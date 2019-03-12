package srv

import (
	"nice-example/api"
	//"nice-example/config"

	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/registry"

	opentracing "github.com/opentracing/opentracing-go"
)

func RunApi(register registry.Registry, tracer opentracing.Tracer, config map[string]interface{}) {
	n := nice.Instance(config["appname"].(string))

	//c := n.LoadConfig("/data/www/golang/src/nice-example/cmd/config.yaml")
	n.SetDI("cache", nice.NewRedis(config["redis"]))
	n.SetDI("db", nice.NewMysql(config["mysql"]))

	api.Router()

	n.Run(config["httpbind"].(string))
}
