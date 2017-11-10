package roomManager

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	broadcastingHost = "localhost"
	broadcastingPort = 6666
	broadcastingURI  = "broadcasting"
	useBroadcasting  = false

	broadcastingConnection websocket.Conn
)

func ConnBroadcasting() {
	conn, err := net.Dial("tcp", broadcastingHost+":"+fmt.Sprint(broadcastingPort))
	if err != nil {
		fmt.Println("[广播站连接失败]：", err)
		return
	}
	u := url.URL{}
	u.Host = broadcastingHost + ":" + fmt.Sprint(broadcastingPort)
	u.Path = broadcastingURI
	u.Scheme = "ws"

	h := http.Header{}
	broadcastingConnection, _, err := websocket.NewClient(conn, &u, h, 1024, 1024)
	if err != nil {
		fmt.Println("[websocket连接失败]：", err)
		return
	}

	//处理读到消息之后
	go func() {
		for {
			_, reader, err := broadcastingConnection.NextReader()
			if err != nil {
				continue
			}
			msg := make([]byte, 1024)
			_, err = reader.Read(msg)
			if err != nil {
				continue
			}
			proMessageFromBroadcast(msg)
		}
	}()
}
