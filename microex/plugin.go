package main

import (
	_ "github.com/kwokhunglee/micro/transport/grpc"
	_ "github.com/kwokhunglee/micro/selector/static"
	_ "github.com/kwokhunglee/micro/registry/etcdv3"
	_ "github.com/kwokhunglee/micro/client/grpc"
	_ "github.com/kwokhunglee/micro/server/grpc"
	_ "github.com/kwokhunglee/micro/broker/rabbitmq"
)
