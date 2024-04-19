package dao

import (
	"github.com/BuyandshipDemo/RpcServer/dao/mysql"
	"github.com/BuyandshipDemo/RpcServer/dao/redis"
)

func init() {
	mysql.Init()
	redis.Init()
}