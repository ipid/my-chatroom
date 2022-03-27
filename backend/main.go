package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hub = NewMessageHub()

func clientMessageHandler(conn *websocket.Conn, clientMsgChan chan<- *ClientMessage) {
	defer func() {
		close(clientMsgChan)
	}()

	for {
		clientMsg := ClientMessage{}
		err := conn.ReadJSON(&clientMsg)
		if err != nil {
			return
		}

		clientMsgChan <- &clientMsg
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// 将历史消息发送给客户端
	for _, msg := range hub.GetHistoryMessages(MAX_MESSAGE_HISTORY) {
		err := conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}

	clientMsgChan := make(chan *ClientMessage, MESSAGE_BUFFER)
	go clientMessageHandler(conn, clientMsgChan)

	hubMsgChan := hub.Subscribe()
	defer hub.Unsubscribe(hubMsgChan)

selectForever:
	for {
		select {
		case clientMsg, ok := <-clientMsgChan:
			if !ok {
				break selectForever
			}

			hub.SendClientMessage(clientMsg, hubMsgChan)

		case hubMsg := <-hubMsgChan:
			err := conn.WriteJSON(hubMsg)
			if err != nil {
				return
			}
		}
	}
}

func main() {
	host := flag.String("host", "", "设置服务所监听的主机名")
	port := flag.Uint("port", 8012, "设置服务所监听的端口号")
	flag.Parse()

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/ws/chat", func(c *gin.Context) {
		handleConnection(c.Writer, c.Request)
	})

	log.Fatal(r.Run(fmt.Sprintf("%s:%d", *host, *port)))
}
