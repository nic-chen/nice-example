package srv

import (
	"nice-example/api"
	"nice-example/constant"

	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/registry"

	opentracing "github.com/opentracing/opentracing-go"
)

func RunApi(register registry.Registry, tracer opentracing.Tracer, config map[string]interface{}) {
	n := nice.Instance(constant.APP_NAME)

	n.SetDI("cache", nice.NewRedis(config["redis"]))
	n.SetDI("db", nice.NewMysql(config["mysql"]))
	n.Conf = config

	api.Router()

	n.Run(config["httpbind"].(string))
}
