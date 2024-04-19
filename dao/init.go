package dao

import (
	"RpcServer/dao/mysql"
	"RpcServer/dao/redis"
)

func init() {
	mysql.Init()
	redis.Init()
}