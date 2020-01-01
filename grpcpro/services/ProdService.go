package services

import (
	"context"
)

type ProdService struct {

}

func(this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {

	  return &ProdResponse{ProdStock:20},nil
}

func (this *ProdService) GetProdStocks(ctx context.Context,size *QuerySize) (*ProdResponseList, error){
	var Prodres []*ProdResponse
	Prodres = make([]*ProdResponse,0,3)
	Prodres = append(Prodres,&ProdResponse{ProdStock:28,})
	Prodres = append(Prodres,&ProdResponse{ProdStock:29,})
	Prodres = append(Prodres,&ProdResponse{ProdStock:30,})
	return &ProdResponseList{
		Prodres:Prodres,
	},nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error){

	ret:=ProdModel{
		ProdId:101,
		ProdName:"测试商品",
		ProdPrice:20.5,
	}
	return &ret,nil
}