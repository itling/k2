package ws

import (
	"admin/service"
	"admin/constant"
	"admin/util"
	"admin/protocol"
	"encoding/base64"
	"io/ioutil"
	"strings"
	"sync"
	"github.com/kingwel-xie/k2/common"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

// 消费kafka里面的消息, 然后直接放入go channel中统一进行消费
func ConsumerKafkaMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	imMsgService := service.ImMessageService{}
	imMsgService.Orm=common.Runtime.GetDb();
	log.Info("start server...")
	for {
		select {
		case conn := <-s.Register:
			log.Info("nuser login userId=",conn.UserId)
			s.Clients[conn.UserId] = conn
			msg := &protocol.Message{
				From:    "",
				To:      "",
				Content: "welcome 职场社交er!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			log.Info("user loginout userId=",conn.UserId)
			if _, ok := s.Clients[conn.UserId]; ok {
				close(conn.Send)
				delete(s.Clients, conn.UserId)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// 一般消息，比如文本消息，视频文件消息等
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
					// 保存消息只会在存在socket的一个端上进行保存，防止分布式部署后，消息重复问题
					_, exits := s.Clients[msg.From]
					if exits {
						saveMessage(&imMsgService,msg)
					}

					if msg.MessageType == constant.MESSAGE_TYPE_USER {
						client, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err == nil {
								client.Send <- msgByte
							}
						}
					} else if msg.MessageType == constant.MESSAGE_TYPE_GROUP {
						//sendGroupMessage(msg, s)
					}
				} else {
					// 语音电话，视频电话等，仅支持单人聊天，不支持群聊
					// 不保存文件，直接进行转发
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}

			} else {
				// 无对应接受人员进行广播
				for _, conn := range s.Clients {
					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.UserId)
					}
				}
			}
		}
	}
}


// 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(imMsgService *service.ImMessageService,message *protocol.Message) {
	// 如果上传的是base64字符串文件，解析文件保存
	if message.ContentType == 2 {
		url := uuid.New().String() + ".png"
		index := strings.Index(message.Content, "base64")
		index += 7

		content := message.Content
		content = content[index:]

		dataBuffer, dataErr := base64.StdEncoding.DecodeString(content)
		if dataErr != nil {
			log.Error("transfer base64 to file error",  dataErr.Error())
			return
		}
		err := ioutil.WriteFile("static/file/"+url, dataBuffer, 0666)
		if err != nil {
			log.Error("write file error",  err.Error())
			return
		}
		message.Url = url
		message.Content = ""
	} else if message.ContentType == 3 {
		// 普通的文件二进制上传
		fileSuffix := util.GetFileType(message.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(message.FileSuffix)
		}
		contentType := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := ioutil.WriteFile("static/file/"+url, message.File, 0666)
		if err != nil {
			log.Error("write file error", err.Error())
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}
	//todo: save db
	imMsgService.SaveMessage(*message)
}
