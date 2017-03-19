package hotgoods

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGetHotGoodsEndpoint(svc UserBaseInfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GerHotGoodsResponse)
		userId := req.UserId
		rsp, err := svc.GetHotGoodsInfo(userId)
		if err != nil{
			return nil, nil
		}
		return rsp, nil
	}
}


func decodeGetHotGoodsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GerHotGoodsResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}



func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type GerUserBaseInfoRequest struct {
	UserId string `json:"user_id"`
}

type GerHotGoodsResponse struct {
	GoodsId string `json:"goods_id"`
	GoodsName string `json:"goods_name"`
	Describe string `json:"describe"`
	Price string `json:"price"`
	Factory string `json:"factory"`
}


