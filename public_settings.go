package roomManager

import (
	"fmt"
	"os"
	"time"
)

var (
	DETAILED_LOG_FLAG           = false                                            //详细日志开关
	NORMAL_LOG_FLAG             = false                                            //常规日志开关
	TRACE_FLAG                  = false                                            //是否打开trace开关
	TRACE_LOG_PATH              = os.ExpandEnv("$GOPATH/trace_logs/trace_log.out") //trace日志地址
	LISTEN_PORT                 = ":8080"                                          //监听端口
	REQUEST_URI                 = "websocket"                                      //请求URI
	CLEAN_TIMER                 = 5 * time.Minute                                  //房间清理定时器
	HALL_TIMEOUT                = 30 * time.Second                                 //大厅房间的连接超过多久之后会被断开
	ROW_LENGTH                  = 1024                                             //单列节点最大长度
	IS_CHECK_USER_WHEN_SEND_MSG = true                                             //发送弹幕时，是否检查用户信息（如果用户ID不存在，则不会发送消息，也不会接收到消息）
)

//设定监听端口号
func SetListenPort(port int) {
	LISTEN_PORT = fmt.Sprint(":", port)
}

//设定接收请求URI
func SetRequestURI(uri string) {
	REQUEST_URI = uri
}

//设定定时清理房间垃圾的间隔时长，单位：分钟
func SetCleanTimes(t int64) {
	CLEAN_TIMER = time.Duration(t) * time.Minute
}

//打开Trace日志开关
func OpenTraceFlag() {
	TRACE_FLAG = true
}

//关闭Trace日志开关
func CloseTraceFlag() {
	TRACE_FLAG = false
}

//打开常规日志开关
func OpenNormalFlag() {
	NORMAL_LOG_FLAG = true
}

//关闭常规日志开关
func CloseNormalFlag() {
	NORMAL_LOG_FLAG = false
}

//打开详细日志开关
func OpenDetailFlag() {
	DETAILED_LOG_FLAG = true
}

//关闭详细日志开关
func CloseDetailFlag() {
	DETAILED_LOG_FLAG = false
}

//设定弹幕发送时是否检查用户身份
func SetCheckUserWhenSendMessage(setting bool) {
	IS_CHECK_USER_WHEN_SEND_MSG = setting
}
