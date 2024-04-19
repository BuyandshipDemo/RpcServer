package main

import (
	_ "RpcServer/dao"
	greet "RpcServer/kitex_gen/example/helloword/greet/itemservice"
	"log"

	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-etcd"
)

func main() {
	// service registry
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	svr := greet.NewServer(new(ItemServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "example.helloworld.greet"}),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

// import (
// 	"log"

// 	greet "RpcServer/kitex_gen/example/helloword/greet/itemservice"
// )

// func main() {
// 	svr := greet.NewServer(new(ItemServiceImpl))

// 	err := svr.Run()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// }