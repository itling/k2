package ws

import (
	"admin/constant"
	"admin/protocol"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserId string
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error("client read message error",err.Error())
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)

		// pong
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err2 := proto.Marshal(pong)
			if nil != err2 {
				log.Error("client marshal message error", err2.Error())
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			//todo 消息队列
			MyServer.Broadcast <- message
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
