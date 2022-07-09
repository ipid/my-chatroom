package main

import (
	"chatroom-backend/def"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var RE_DICE = regexp.MustCompile(`^\s*(\d{1,2})d(\d{1,2})\s*$`)

var hub = NewMessageHub()

func clientRequestHandler(conn *websocket.Conn, chanClientRequest chan<- *ClientRequest) {
	defer close(chanClientRequest)

	// 将历史消息发送给客户端
	for index, msg := range hub.GetHistoryMessages(MAX_MESSAGE_HISTORY) {
		log.Printf("Sending history message %d:\n", index)

		err := conn.WriteJSON(ServerResponse{
			Type: RESP_TYPE_NEW_MESSAGE,
			Data: msg,
		})
		if err != nil {
			log.Printf("Failed to send history message: %v\n", err.Error())
			return
		}
	}

	for {
		unknownReq := UnknownClientRequest{}
		if err := conn.ReadJSON(&unknownReq); err != nil {
			log.Printf("Parse request as UnknownClientRequest failed: %v\n", err.Error())
			return
		}

		var req *ClientRequest
		if unknownReq.Type == REQ_TYPE_MESSAGE {
			req = &ClientRequest{
				Type:       unknownReq.Type,
				MsgRequest: &ClientRequestMessage{},
			}
			if err := json.Unmarshal(unknownReq.Data, req.MsgRequest); err != nil {
				log.Printf("Client sent a `send-msg` request, but request data can't be parsed due to: %v\n", err.Error())
				return
			}
			unknownReq.Data = nil
		} else {
			log.Printf("Unknown client request type: <%v>\n", unknownReq.Type)
			return
		}

		chanClientRequest <- req
	}
}

func handleClientRequest(req *ClientRequest, myId int64) *ServerResponse {
	msg := &def.Message{
		Id:          -1,
		Sender:      req.MsgRequest.Sender,
		SenderId:    myId,
		TimestampMs: time.Now().UnixMilli(),
	}

	match := RE_DICE.FindStringSubmatch(req.MsgRequest.Content)

	if match != nil {
		dNum, _ := strconv.Atoi(match[1])
		dMax, _ := strconv.Atoi(match[2])

		if 1 <= dNum && dNum <= 99 && 1 <= dMax && dMax <= 99 {
			msg.Dice = &MessageDataDice{
				Dices: util.getDices(dNum, dMax),
			}
		}
	}

	if msg.Dice == nil {
		msg.RichText = &MessageDataRichText{
			Content: req.MsgRequest.Content,
		}
	}

	hub.SubmitAndBroadcastMessage(msg)

	return &ServerResponse{
		Type: RESP_TYPE_MESSAGE_ACKNOWLEDGE,
		Data: msg,
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// 开启一个用于读取用户请求的 Goroutine
	chanClientReq := make(chan *ClientRequest)
	go clientRequestHandler(conn, chanClientReq)

	// 本连接对应的 id，需要填进 Message 里
	chanWriteToClient, myId := hub.Subscribe()
	log.Printf("[%d] New connection\n", myId)

selectForever:
	for {
		select {
		case req, ok := <-chanClientReq:
			if !ok {
				log.Printf("[%d] chanClientReq closed\n", myId)
				break selectForever
			}

			log.Printf("[%d] received a client request: %v, \n", myId, req)
			result := handleClientRequest(req, myId)

			if err := conn.WriteJSON(result); err != nil {
				log.Printf("[%d] Failed to write ack message due to: %v\n", myId, err.Error())
				break selectForever
			}
		case resp := <-chanWriteToClient:
			if err := conn.WriteJSON(resp); err != nil {
				log.Printf("[%d] Failed to write message from other senders due to: %v\n", myId, err.Error())
				break selectForever
			}
		}
	}

	hub.Unsubscribe(myId)
	log.Printf("[%d] Connection exit\n", myId)
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
