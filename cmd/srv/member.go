package srv

import (
	"fmt"
	"net"

	"nice-example/proto/member"
	"nice-example/services/membersrv"
	"nice-example/config"

	"github.com/nic-chen/nice/micro"
	"github.com/nic-chen/nice/micro/registry"

	opentracing "github.com/opentracing/opentracing-go"
)

func RunMemberSrv(register registry.Registry, tracer opentracing.Tracer) {
	var (
		err    error
	)

	service := membersrv.NewMemberService()

	listen := net.JoinHostPort(config.SrvHost, config.SrvPort)

	var opts []micro.Option
	if register!=nil {
		opts = append(opts, micro.WithRegistry(register, config.MemberSrvName, listen))
	}
	if tracer!=nil {
		opts = append(opts, micro.WithTracer(tracer))
	}

	server, err := micro.NewServer(config.MemberSrvName, opts...)

	if err != nil {
		panic(fmt.Errorf("%s server start error:%s", config.MemberSrvName, err))
	}

	rpc := server.BuildGrpcServer()
	member.RegisterMemberServer(rpc, service)
	
	err = server.Run(rpc, listen); 
}
