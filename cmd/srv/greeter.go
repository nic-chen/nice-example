package srv

import (
	"fmt"
	"net"
	"github.com/nic-chen/nice/micro"
	"github.com/nic-chen/nice/micro/registry"
	"../../proto/greeter"
	"../../services/greetersrv"
	"../../config"

)

func RunGreeter(register registry.Registry) {
	var (
		err    error
	)

	service := greetersrv.NewGreeterService()

	listen := net.JoinHostPort(config.SrvHost, config.SrvPort)

	var opts = []micro.Option{
		// micro.WithTracer(tracer),
		// micro.WithLogger(dlog.Strict()),
		micro.WithRegistry(register, config.SrvName, listen),
	}


	server, err := micro.NewServer(config.SrvName, opts...)

	if err != nil {
		panic(fmt.Errorf("%s server start error:%s", config.SrvName, err))
	}

	rpc := server.BuildGrpcServer()
	greeter.RegisterGreeterServer(rpc, service)
	
	err = server.Run(rpc, listen); 
}
