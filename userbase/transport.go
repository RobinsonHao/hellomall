package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGetUserBaseInfoEndpoint(svc UserBaseInfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GerUserBaseInfoRequest)
		userId := req.UserId
		rsp, err := svc.GetUserBaseInfo(userId)
		if err != nil{
			return nil, nil
		}
		return rsp, nil
	}
}


func decodeGetUserBaseInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GerUserBaseInfoRequest
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

type GerUserBaseInfoResponse struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserTag string `json:"user_tag"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
	Email string `json:"emmail"`
}


