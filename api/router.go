package api

import (
	"github.com/nic-chen/nice"
	"../config"
	"./controller"
)

func Router() {
	n := nice.Instance(config.APP_NAME)
	
	n.Get("/member/:id", controller.Member.Info);
	n.Get("/greeter/:name", controller.Greeter.Hello);
	n.Get("/info/:id", controller.Member.Basic);
}