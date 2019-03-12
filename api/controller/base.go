package controller

import (
	"context"
	"time"

	"nice-example/constant"

	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/tracing"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/nic-chen/nice/micro/registry/etcdv3"
	"google.golang.org/grpc"
)

type Controller struct {
	Name string
}

type JsonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func New(name string) *Controller {
	ctl := &Controller{
		Name: name,
	}

	return ctl
}

func RenderJson(c *nice.Context, code int, message string, data interface{}) {
	ret := new(JsonResp)
	ret.Code = code
	ret.Message = message
	if data != nil {
		ret.Data = data
	}
	c.JSON(200, ret)
}

func newSrvDialer(serviceName string) *grpc.ClientConn {
	var err error
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	n := nice.Instance(constant.APP_NAME)

	//如果配置了jaeger
	if len(n.Conf["jaeger"].(string)) > 0 {
		tracer, err := tracing.Init(n.Conf["clientname"].(string), n.Conf["jaeger"].(string))
		if err == nil {
			opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var conn *grpc.ClientConn

	//如果配置了etcd
	if len(n.Conf["etcd"].(string)) > 0 {
		r := etcdv3.NewResolver(serviceName)
		b := grpc.RoundRobin(r)
		opts = append(opts, grpc.WithBalancer(b))

		conn, err = grpc.DialContext(ctx, n.Conf["etcd"].(string), opts...)
	} else {
		conn, err = grpc.DialContext(ctx, serviceName, opts...)
	}

	cancel()

	if err != nil {
		panic(err)
	}

	return conn
}
