package api

import (
	"github.com/nic-chen/nice"
	"nice-example/config"
	"nice-example/api/controller"
)

func Router() {
	n := nice.Instance(config.APP_NAME)
	
	n.Get("/member/:id", controller.Member.Info);
	n.Get("/info/:id", controller.Member.Basic);
}