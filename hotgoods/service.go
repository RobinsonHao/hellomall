package hotgoods

import (
	"errors"
	//"strings"
)


type HotGoodsService interface {
	GetHotGoodsInfo(string) (interface{}, error)
}

type hotGoogdsService struct{}

func (hotGoogdsService) GetHotGoods() (interface{}, error) {
	var rsp  GerHotGoodsResponse
		rsp.GoodsId = "1001"
		rsp.GoodsName = "abc"
		rsp.Describe = "goods"
		rsp.Factory = "china"
		rsp.Price = "100"
	return rsp, nil
}


// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
