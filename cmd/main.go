package main

import (
	"flag"
	"fmt"
	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/registry"
	_ "github.com/nic-chen/nice/micro/registry/etcdv3"
	"github.com/nic-chen/nice/micro/tracing"
	opentracing "github.com/opentracing/opentracing-go"
	"log"
	"nice-example/cmd/srv"
	"os"
	"os/signal"
	//"strings"
	"syscall"
)

var conf string

func usage() {
	fmt.Fprintf(os.Stderr, "请用 -c 指定配置文件 \n")
}

func init() {
	flag.StringVar(&conf, "c", "/data/config.yaml", "配置文件绝对路径")
}

func main() {

	var (
		register registry.Registry
		tracer   opentracing.Tracer
		err      error
	)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	flag.Parse()

	config := nice.LoadConfig(conf)

	if config["appname"].(string) != "" && config["serverhost"].(string) != "" && config["serverport"].(string) != "" && config["etcd"].(string) != "" {
		options := &registry.Options{
			Name: config["appname"].(string),
			Host: config["serverhost"].(string),
			Port: config["serverport"].(string),
			TTL:  config["servercheckttl"].(int),
			Ssrv: config["etcd"].(string),
		}
		register, err = registry.DefaultRegistry(options)
		log.Printf("NamingAddr: %s", config["etcd"])
		if err != nil {
			panic(err)
		}
	}

	if config["appname"].(string) != "" && config["jaeger"].(string) != "" {
		tracer, err = tracing.Init(config["appname"].(string), config["jaeger"].(string))
		if err != nil {
			panic(err)
		}
	}

	//监听退出
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		<-ch
		register.UnRegister()
		os.Exit(1)
	}()

	srv.RunAll(register, tracer, config)
}
