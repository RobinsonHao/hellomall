package main

import (
	"encoding/json"
	"log"

	"hellomall/userbase/srv/proto"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"

)

// 
type TrunkData struct{
	ctx *context.Context
	req *hellomall.Request
}

type WorkProcessOper interface {
	GetWorkRequestPre
	ProcessWorkRequest
	ProcessWorkResponse
}

type WorkRequest struct {
	ctx *context.Context
	req interface 
	preDataName string 
} 

type WorkData struct{
	// 用于存储所有请求的列表
    var WorkRequestMap = make(map[string] *WorkRequet)

    // 用于存储所有应答的列表
    var WorkResponseMap = make(map[string] interface)

    // 最终给用户返回的数据
    rsp *hellomall.Response
}

// 存储工作单元是否被调用过
var WorkItemStatus = make(map[string] int32)

// 存储工作单元对应的结构
var WorkItemStruct = make(map[string] interface)




// 本方法是获取首页数据的入口
func (s *BaseData) GetMallHomeInfo(ctx context.Context, req *mallhome.Request, rsp *mallhome.Response) error {
	userId := req.UserId

	if userId < 1 {
        rsp.errno = ERR_PARAM_ERROR
        rsp.errmsg = Err_Msg[rsp.errno]
	    log.Fatalf("input param error, the userId: %d", userId)  
        return nil
	}

    workItemTotalNum = len(WorkItemStatus)

    // 建立一个异步接收通道
	done := make(chan map[string] *, workItemTotalNum)
	var wg sync.WaitGroup

    // 遍历所有工作单元请元，直到处理完或者超时为止
	for {
		for _, workItemName := range WorkItemNames {
    	    workItemStat := WorkItemName[workItemName]

    	    // 遍历没有发送过淤请求的工作单元
            if 0 == workItemStat {
	            preRet,err := WorkItemStruct[workItemName].GetWorkRequestPre(req, nil, &WorkData)

                // 发送没有前置数据或者前置数据已具备的数据
	            if (nil == preRet || nil != WorkData.WorkResponseMap[workItemName]) {
	            	workReqCli, err := WorkItemStruct.ProcessWorkRequest(req, nil, &WorkData)
	                WorkItemStatus[workItemName] = 1
	            	go func () {
	            		wg.Add(1)
	            		workItemRsp, err := workReqCli.Endpoint()(ctx, struct{}{})
	            		workRsp := map[string]string{WorkItemName: workItemRsp}
	            		done <- workRsp
	            	}
	            }

	            time.Sleep(time.Millisecond * 100)
	    	 }
    	   
        }
        // 解析工作单元的输出数据
        go func() {
        	for _, rsp := range done {
	    	    append(WorkData.WorkResponseMap,rsp)
	    	 	err := WorkItemStruct.ProcessWorkResponse(req, nil, &WorkData)
	    	 	wg.Done()
            }
        }

        rspLen = len(WorkData.WorkResponseMap)

        if rspLen == workItemTotalNum{
        	break
        }


	}

    limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(3000, 10))
	homemallEndpoint = limiter(homemallEndpoint)

    
    rsp.errno = ERR_SUCCESS
    rsp.errmsg = Err_Msg[rsp.errno]
     
	return nil
}

func init(){
	WorkItemStatus['UserBaseInfo'] = 0 // 用户基础信息
	WorkItemStatus['UserTag'] = 0      // 用户标签
	WorkItemStatus['UserRecGoods'] = 0  // 用户推荐端口
	WorkItemStatus['TopFeed'] = 0 // 顶部feed流信息
	WorkItemStatus['HotGoods'] = 0 // 热门商品

	WorkItemStruct['UserBaseInfo'] = UserBaseInfo
	WorkItemStruct['UserTag'] = UserTag
	WorkItemStruct['UserRecGoods'] = UserRecGoods
	WorkItemStruct['TopFeed'] = TopFeed
	WorkItemStruct['HotGoods'] = HotGoods
}

