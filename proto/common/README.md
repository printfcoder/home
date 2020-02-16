# COMPILE

Be sure all imported proto files are in $GOPATH

## ERROR

```bash
$ protoc --proto_path=. --go_out=.  error/error.proto
```

## RESPONSE

```bash
$ protoc --proto_path=.:$GOPATH/src --go_out=.  response/response.proto
```

