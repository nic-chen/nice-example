package srv

import (
	"fmt"
	"net"

	"../../proto/greeter"
	"../../services/greetersrv"
	"../../config"

	"github.com/nic-chen/nice/micro"
	"github.com/nic-chen/nice/micro/registry"

	opentracing "github.com/opentracing/opentracing-go"
)

func RunGreeter(register registry.Registry, tracer opentracing.Tracer) {
	var (
		err    error
	)

	service := greetersrv.NewGreeterService()

	listen := net.JoinHostPort(config.SrvHost, config.SrvPort)

	var opts = []micro.Option{
		// micro.WithTracer(tracer),
		// micro.WithLogger(dlog.Strict()),
		micro.WithRegistry(register, config.SrvName, listen),
		micro.WithTracer(tracer),
	}


	server, err := micro.NewServer(config.SrvName, opts...)

	if err != nil {
		panic(fmt.Errorf("%s server start error:%s", config.SrvName, err))
	}

	rpc := server.BuildGrpcServer()
	greeter.RegisterGreeterServer(rpc, service)
	
	err = server.Run(rpc, listen); 
}
