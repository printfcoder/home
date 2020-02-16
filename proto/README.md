# COMPILE

Be sure all imported proto files are in $GOPATH

## error

```bash
$ protoc --proto_path=. --go_out=.  common/error/error.proto
```

## response

```bash
$ protoc --proto_path=.:$GOPATH/src --go_out=.  common/response/response.proto
```

## time

```bash
$ protoc --proto_path=.:$GOPATH/src --go_out=.  common/time/time.proto
```

## finance/book

```bash
$ protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. finance/book/book.proto
```