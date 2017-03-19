package main

import (
	"errors"
	//"strings"
)


type UserBaseInfoService interface {
	GetUserBaseInfo(string) (interface{}, error)
}

type userBaseInfoService struct{}

func (userBaseInfoService) GetUserBaseInfo(s string) (interface{}, error) {
	var rsp  GerUserBaseInfoResponse
		rsp.UserId = "1001"
		rsp.UserName = "abc"
		rsp.UserTag = "tv"
		rsp.Sex = "man"
		rsp.Phone = "1390000001"
		rsp.Email = "test@qq.com"
	return rsp, nil
}


// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
