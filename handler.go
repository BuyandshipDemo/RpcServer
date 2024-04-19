package main

import (
	"RpcServer/dao/mysql"
	"RpcServer/dao/redis"
	greet "RpcServer/kitex_gen/example/helloword/greet"
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *greet.GetItemReq) (resp *greet.GetItemResp, err error) {
	// TODO: Your code here...
	klog.CtxDebugf(ctx, "req: %+v", req)
	resp = greet.NewGetItemResp()
	msg, err := mysql.Query(req.GetId())
	resp.Item = greet.NewItem()
	if err != nil {
		resp.Item.Msg = err.Error()
		return
	}
	resp.Item.Id = req.GetId()
	resp.Item.Msg = msg
	// set redis 
	res := redis.RedisCli.Set(ctx, "helloworld", resp.Item.Msg, 0)
	v, err := res.Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(v)
	return
}
