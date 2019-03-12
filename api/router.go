package api

import (
	"github.com/nic-chen/nice"
	"nice-example/api/controller"
	"nice-example/constant"
)

func Router() {
	n := nice.Instance(constant.APP_NAME)
	n.Get("/member/:id", controller.Member.Info)
	n.Get("/info/:id", controller.Member.Basic)
}
