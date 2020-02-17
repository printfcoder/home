module github.com/printfcoder/home/finance

go 1.13

replace github.com/printfcoder/home/proto v0.0.0 => /Users/shuxian/workspace/go/src/github.com/printfcoder/home/proto

require (
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/micro/go-micro/v2 v2.1.0 // indirect
	github.com/printfcoder/home/proto v0.0.0
)
