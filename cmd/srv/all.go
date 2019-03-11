package srv

import (
	"github.com/nic-chen/nice/micro/registry"
	opentracing "github.com/opentracing/opentracing-go"
)

func RunAll(register registry.Registry, tracer opentracing.Tracer, config map[string]interface{}) {
	go RunMemberSrv(register, tracer, config)
	RunApi(register, tracer, config)
}
