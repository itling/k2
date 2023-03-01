package ws

import (
	"strconv"
	"net/http"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var log = logger.Logger("ws")



func RunSocekt(c *gin.Context) {
	identity := service.GetIdentity(c)

	log.Infof("newUser %s", identity.Username)

	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		UserId: strconv.Itoa(identity.UserId),
		Conn: ws,
		Send: make(chan []byte),
	}

	MyServer.Register <- client
	go client.Read()
	go client.Write()
}
