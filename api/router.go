package api

import (
	"github.com/nic-chen/nice"
	"log"
	"nice-example/api/controller"
	"nice-example/config"
)

func Router() {
	n := nice.Instance(config.APP_NAME)

	log.Printf("router")

	n.Get("/member/:id", controller.Member.Info)
	n.Get("/info/:id", controller.Member.Basic)
}
