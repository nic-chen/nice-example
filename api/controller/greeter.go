package controller

import (
	"log"
	//"encoding/json"
	"context"
	"github.com/nic-chen/nice"
	//"../../config"
	proto "../../proto/greeter"

	"../../config"
)

type greeter struct{}
var Greeter = greeter{}

func (greeter) Hello(c *nice.Context) {
	name := c.Param("name");

	//n := nice.Instance(config.APP_NAME);

    //服务名
	conn := newSrvDialer(config.TestSrvName)
    //grpc client
	client := proto.NewGreeterClient(conn);

	req := &proto.Request{
		Name: name,
	}

	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Panicf("dialer error: %v", err)
	}

	RenderJson(c, 0, "", res)
}
