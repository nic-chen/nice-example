package main

import (
	"nice"
	"test/config"
	"test/controller"
)

func Router() {
	n := nice.Instance(config.APP_NAME)
	
	n.Get("/member/:id", controller.Member.Info);
}