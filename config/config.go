//config

package config

import (
	//"time"
)

var (
	APP_NAME  = "nice";
	PPROFBIND = []string{"localhost:2045"}
	HTTPBIND  = "0.0.0.0:8202"
)

const (
	Debug            = true;

	MYSQLHOST        = "127.0.0.1";
	MYSQLDB          = "carrier";
	MYSQLUSER        = "root";
	MYSQLPWD         = "";

	DBCHARSET        = "utf8";
	DBCONNOPEN       = 100;
	DBCONNIDLE       = 10;

	REDISHOST        = "127.0.0.1:6379";
	REDISDB          = 0;
	REDISPWD         = "";

)