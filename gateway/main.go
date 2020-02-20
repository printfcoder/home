// Package api is an API Gateway
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/v2"
	ahandler "github.com/micro/go-micro/v2/api/handler"
	arpc "github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/api/resolver"
	rrmicro "github.com/micro/go-micro/v2/api/resolver/micro"
	"github.com/micro/go-micro/v2/api/router"
	regRouter "github.com/micro/go-micro/v2/api/router/registry"
	"github.com/micro/go-micro/v2/api/server"
	httpapi "github.com/micro/go-micro/v2/api/server/http"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	Name         = "go.micro.api"
	Address      = ":8080"
	Handler      = "meta"
	Resolver     = "micro"
	RPCPath      = "/rpc"
	APIPath      = "/"
	ProxyPath    = "/{service:[a-zA-Z0-9]+}"
	Namespace    = "go.micro.api"
	HeaderPrefix = "X-Micro-"
	EnableRPC    = false
)

func run(ctx *cli.Context, srvOpts ...micro.Option) {
	// Init API
	var opts []server.Option

	// create the router
	var h http.Handler
	r := http.NewServeMux()
	h = r

	// return version and list of services
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf(`{"version": "%s"}`, ctx.App.Version)
		w.Write([]byte(response))
	})

	srvOpts = append(srvOpts, micro.Name(Name))
	if i := time.Duration(ctx.Int("register_ttl")); i > 0 {
		srvOpts = append(srvOpts, micro.RegisterTTL(i*time.Second))
	}
	if i := time.Duration(ctx.Int("register_interval")); i > 0 {
		srvOpts = append(srvOpts, micro.RegisterInterval(i*time.Second))
	}

	// initialise service
	service := micro.NewService(srvOpts...)

	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(Namespace),
		resolver.WithHandler(Handler),
	}

	// default resolver
	rr := rrmicro.NewResolver(ropts...)

	rt := regRouter.NewRouter(
		router.WithNamespace(Namespace),
		router.WithHandler(arpc.Handler),
		router.WithResolver(rr),
		router.WithRegistry(service.Options().Registry),
	)

	rp := arpc.NewHandler(
		ahandler.WithNamespace(Namespace),
		ahandler.WithRouter(rt),
		ahandler.WithService(service),
	)

	r.Handle(APIPath, rp)

	// create the server
	api := httpapi.NewServer(Address)
	api.Init(opts...)
	api.Handle("/", h)

	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := api.Stop(); err != nil {
		log.Fatal(err)
	}
}
