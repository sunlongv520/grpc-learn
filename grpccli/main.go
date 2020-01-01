package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"gprccli/services"
	"io"
	"log"
	"time"
)

func main2(){

	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient:=services.NewProdServiceClient(conn)
	ctx:=context.Background()

	//prodRes,err:=prodClient.GetProdStock(context.Background(),
	//	&services.ProdRequest{ProdId:12})
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(prodRes.ProdStock)

	//response,err := prodClient.GetProdStocks(ctx,&services.QuerySize{Size:10})
	//
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(response.Prodres[2].ProdStock)

	prod,err :=prodClient.GetProdInfo(ctx,&services.ProdRequest{ProdId:12})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(prod)
}

func main3(){
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	ctx:=context.Background()
	t:=timestamp.Timestamp{Seconds:time.Now().Unix()}
	orderClient:=services.NewOrderSerivceClient(conn)
	res,_:= orderClient.NewOrder(ctx,&services.OrderMain{
		OrderId:1001,
		OrderNo:"20190809",
		OrderMoney:90,
		OrderTime:&t,
	})
	fmt.Println(res)
}

func main4(){
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	ctx:=context.Background()
	userClient:=services.NewUserServiceClient(conn)
	var i int32
	req:=services.UserScoreRequest{}
	req.Users=make([]*services.UserInfo,0)

	for i=1;i<20;i++{
		req.Users=append(req.Users,&services.UserInfo{UserId:i})
	}
	res,_ := userClient.GetUserScore(ctx,&req)
	fmt.Println(res.Users)

}

//服务端流
func main5()  {
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	ctx:=context.Background()

	userClient:=services.NewUserServiceClient(conn)
	var i int32
	req:=services.UserScoreRequest{}
	req.Users=make([]*services.UserInfo,0)

	for i=1;i<6;i++{  //加了5条用户信息
		req.Users=append(req.Users,&services.UserInfo{UserId:i})
	}

	stream,err:=userClient.GetUserScoreByServerStream(ctx,&req)
	if err!=nil{
		log.Fatal(err)
	}
	for{
		res,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(res.Users)

	}
}

//客户端流
func main6(){
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	ctx:=context.Background()
	userClient:=services.NewUserServiceClient(conn)
	var i int32
	if err!=nil{
		log.Fatal(err)
	}
	stream,err:=userClient.GetUserScoreByClientStream(ctx)
	if err!=nil{
		log.Fatal(err)
	}

	for j:=1;j<=3;j++{
		req:=services.UserScoreRequest{}
		req.Users=make([]*services.UserInfo,0)
		for i=1;i<=5;i++{  //加了5条用户信息  假设是一个耗时的过程
			req.Users=append(req.Users,&services.UserInfo{UserId:i})
		}
		err:=stream.Send(&req)
		if err!=nil{
			log.Println(err)
		}
	}
	res,_:=stream.CloseAndRecv()
	fmt.Println(res.Users)
}

//双向流
func main(){
	conn,err:=grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	ctx:=context.Background()
	userClient:=services.NewUserServiceClient(conn)
	var i int32
	if err!=nil{
		log.Fatal(err)
	}
	stream,err:=userClient.GetUserScoreByTWS(ctx)
	if err!=nil{
		log.Fatal(err)
	}

	var uid int32=1
	for j:=1;j<=3;j++{
		req:=services.UserScoreRequest{}
		req.Users=make([]*services.UserInfo,0)
		for i=1;i<=5;i++{  //加5条用户信息  假设是一个耗时的过程
			req.Users=append(req.Users,&services.UserInfo{UserId:uid})
			uid++
		}
		err:=stream.Send(&req)
		if err!=nil{
			log.Println(err)
		}
		res,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Println(err)
		}
		fmt.Println(res.Users)

	}
}
