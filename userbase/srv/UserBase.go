package main

import (
	"encoding/json"
	"log"

	"hellomall/userbase/srv/proto"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"

	"github.com/Go-SQL-Driver/Mysql"
	"database/sql"
)

type BaseData struct {}

// get use base info return user base info  for requested user id
func (s *BaseData) GetUserBaseInfo(ctx context.Context, req *userbase.Request, rsp *userbase.Response) error {
	userId := req.UserId

	if userId < 1 {
        rsp.errno = ERR_PARAM_ERROR
        rsp.errmsg = Err_Msg[rsp.errno]
	    log.Fatalf("input param error, the userId: %d", userId)  
        return nil
	}

	userBaseInfo, err :=  GetUserBaseInfoFrmDb(userId)

	if (nil != err){
		rsp.errno = ERR_CALL_DB_FAIL
        rsp.errno = Err_Msg[rsp.errno]
        log.Fatalf("get user base info from db fail, %v", err)
        return nil
	}

	rsp.errno = ERR_SUCCESS
	rsp.errmsg = Err_Msg[rsp.errno]
	rsp.userBaseInfo = userBaseInfo

	return nil
}

// get user base info frm db
func GetUserBaseInfoFrmDb(userId int64) (userbase.UserBaseInfo, error){
    db, err := sql.Open("mysql", "user:password@test?charset=utf8")
    
    if (err != nil) {
    	log.Fatalf("init db fail")
    	return nil, errors.New('init db fail')
    }

    sql := "selelct user_id, user_name, sex, real_user_name, phone, email from user_base_info where user_id= %d"
    sql = fmp.Spintf(sql, userId)
    rows, err = db.Query(sql)

    if (nil == rows[0] && nil !== err){
    	return nil, nil
    }

    var userId int64
    var userName string
    var sex int32
    var realUserName string
    var phone string
    var email string 

    err = rows.Scan(&userId, &userName, &sex, &realUserName, &phone, &email)

    if (nil != err) {
        log.Fatalf('query db fail, the sql=%s', sql)
        return nil, errors.New('quer db fail')
    } 

    userbase.UserBaseInfo.UserId = userId
    userbase.UserBaseInfo.UserName = userName
    userbase.UserBaseInfo.Sex = sex
    userbbse.UserBaseInfo.RealUserName = realUserName
    userbase.UserBaseInfo.Phone = phone
    userbase.UserBaseInfo.Email = email

    return &userbase.UserBaseInfo, nil
}




func main() {
	service := micro.NewService(
		micro.Name("hellomall.srv.userbase"),
	)

	service.Init()

	userbase.RegisterUserBaseHandler(service.Server(), new(BaseData))

	service.Run()
}
