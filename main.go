package main

import (
	_ "github.com/BuyandshipDemo/RpcServer/dao"
	greet "github.com/BuyandshipDemo/RpcServer/kitex_gen/example/helloword/greet/itemservice"
	"log"

	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-etcd"
	cc "github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
)

func main() {
	// service registry
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	// ConfigCenter
	etcdClient, _ := cc.NewClient(cc.Options{
		Node: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	svr := greet.NewServer(new(ItemServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "example.helloworld.greet"}),
		server.WithRegistry(r),
		server.WithSuite(etcdServer.NewSuite("example.helloworld.greet", etcdClient)),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
