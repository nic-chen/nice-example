package srv

import (
	"fmt"
	"net"

	//"nice-example/config"
	"nice-example/proto/member"
	"nice-example/services/membersrv"

	"github.com/nic-chen/nice/micro"
	"github.com/nic-chen/nice/micro/registry"

	opentracing "github.com/opentracing/opentracing-go"
)

func RunMemberSrv(register registry.Registry, tracer opentracing.Tracer, config map[string]interface{}) {
	var (
		err error
	)

	service := membersrv.NewMemberService()

	listen := net.JoinHostPort(config["serverhost"].(string), config["serverport"].(string))

	var opts []micro.Option
	if register != nil {
		opts = append(opts, micro.WithRegistry(register, config["servername"].(string), listen))
	}
	if tracer != nil {
		opts = append(opts, micro.WithTracer(tracer))
	}

	server, err := micro.NewServer(config["servername"].(string), opts...)

	if err != nil {
		panic(fmt.Errorf("%s server start error:%s", config["servername"].(string), err))
	}

	rpc := server.BuildGrpcServer()
	member.RegisterMemberServer(rpc, service)

	err = server.Run(rpc, listen)
}
