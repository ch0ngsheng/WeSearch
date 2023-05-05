# cmd

```shell
goctl rpc protoc ./rpc/userdoc.proto --go_out=. --go-grpc_out=. --zrpc_out=./rpc --home ../../deploy/goctl/
```

```shell
goctl docker -go rpc/userdoc.go
```

```shell
# create Kafka topic

```

```shell
# model
goctl model mysql ddl -src="./model/sql/model.sql" -dir="./model" --home ../../deploy/goctl/
```