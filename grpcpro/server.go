package main

import (
"google.golang.org/grpc"
"grpcpro/services"
"net"
)

func main()  {
	rpcServer:=grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))//商品服务
	services.RegisterOrderSerivceServer(rpcServer,new(services.OrdersService))//订单服务
	services.RegisterUserServiceServer(rpcServer,new(services.UserService))

	lis,_:=net.Listen("tcp",":8081")

	rpcServer.Serve(lis)


}
