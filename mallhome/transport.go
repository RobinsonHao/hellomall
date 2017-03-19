package mallhome

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeMallHomeEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(mallHomeRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}



func decodeMallHomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request mallHomeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}



func decodeGetUserInfoResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response GerUserBaseInfoResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type mallHomeRequest struct {
	UserId string `json:"user_id"`
}

type mallHomeResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type GerUserBaseInfoResponse struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserTag string `json:"user_tag"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
	Email string `json:"emmail"`
}

