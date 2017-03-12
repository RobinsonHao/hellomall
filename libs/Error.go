package hello_mall_libs

const ERR_SUCCESS = 0
const ERR_PARAM_ERROR = 1
const ERR_CALL_DB_FAIL = 2
const ERR_CALL_RPC_FAIL = 3
const ERR_USE_IS_NOT_EXIST = 4

var Err_Msg = make(map[int64]string)
Err_Msg[ERR_SUCCESS] = 'success'
Err_Msg[ERR_PARAM_ERROR] = 'param error'
Err_Msg[ERR_CALL_DB_FAIL] = 'call db fail'
Err_Msg[ERR_CALL_RPC_FAIL] = 'call rpc fail'
Err_Msg[ERR_USE_IS_NOT_EXIST] = 'user is not exist'