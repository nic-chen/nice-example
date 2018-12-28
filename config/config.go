//config

package config

import (
	//"time"
)

var (
	APP_NAME    = "nice"
	//API绑定IP端口
	HTTPBIND    = "0.0.0.0:8202"
	//etcd地址
	NamingAddr  = ""     //http://127.0.0.1:2379 留空不使用etcd
	//微服务相关配置
	SrvName     = "localhost:50001" //"member-srv"  //不使用服务发现时，服务名改成host:port即可
	SrvHost     = "localhost"
	SrvPort     = "50001"
	SrvCheckTTL = 30
	CliName     = "member-cli"
	MemberSrvName = SrvName

	//跟踪服务jaeger地址
	JaegerAddr  = ""         //127.0.0.1:6831   留空不使用jaeger
)

const (
	Debug      = true
	
	//mysql配置
	MYSQLHOST  = "127.0.0.1"
	MYSQLDB    = "carrier"
	MYSQLUSER  = "root"
	MYSQLPWD   = ""
	DBCHARSET  = "utf8"
	DBCONNOPEN = 100
	DBCONNIDLE = 10
	
	//redis配置
	REDISHOST  = "127.0.0.1:6379"
	REDISDB    = 0
	REDISPWD   = ""
)