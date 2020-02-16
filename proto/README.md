# COMPILE

Be sure all imported proto files are in $GOPATH

## common

```bash
protoc --proto_path=. --go_out=.  common/error/error.proto
protoc --proto_path=.:$GOPATH/src --go_out=.  common/time/time.proto
protoc --proto_path=.:$GOPATH/src --go_out=.  common/response/response.proto
```

## finance/book

```bash
$ protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. finance/book/book.proto
```