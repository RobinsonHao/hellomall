package main

import (
	"encoding/json"
	"log"

	"hellomall/userbase/srv/proto"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"

    "encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	
)

type UserBaseInfoData struct{
	WorkProcessOper
}

// 设置前置依赖数据名称，如果没有，则返回空，如果有，则返回依赖的数据名, mediaData 是用于存于中间数据，为了类型方便使用，因此使用了
// JSON字符串，这个类型在通常情况下不会使用
func GetWorkRequestPre(userReq *hellomall.Request, mdeiaData string, workData *WorkData) string, error{
	return nil, nil
}

// 本方法的目的是根据用户输入参数或者已有的中间数据构造工作所需的请求
func ProcessWorkRequest(userReq *hellomall.Request, mediaData string, workData *WorkData) interface, error{
    cli := http.NewClient(
         'post'
         'http://localhost:8080/userbase/GetUserBaseInfo'
        EncodeRequestFuncForUserBaseInfo
        ProcessWorkResponse
    	)
     
    return cli, nil
}

// 本方法是对工作作数据请求进行编码
func EncodeRequestFuncForUserBaseInfo(userReq *hellomall.Request, mdeiaData string, workData *WorkData) *userbase.Request, error{
    userbase.Request.userId = hellomall.Request.userId
    return &userbse.Request
}

// 本方法是用于处理本工作返回的RPC结果
func ProcessWorkResponse(userReq *hellomall.Request, mediaData string, workData *WorkData) error{
    userBaseRequestTmp = workData.WorkResponseMap['UserBaseInfoData']
    userBaseRequest = decode(userBaseRequestTmp) // 伪代码，对userbase的返回值进行解码
    workData.rsp.userBaseInfo = userBaseRequest
    return nil
}







    


