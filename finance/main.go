package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/printfcoder/home/finance/book"
	"github.com/printfcoder/home/finance/book/handler"
	bProto "github.com/printfcoder/home/proto/finance/book"
)

func main() {
	s := micro.NewService(
		micro.Name("home.srv.fin"),
	)

	s.Init(micro.Action(func(context *cli.Context) error {
		book.Init()
		return nil
	}))

	err := bProto.RegisterBookHandler(s.Server(), new(handler.BookHandler))
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
