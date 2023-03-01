package dto

type MessageRequest struct {
	MessageType    int32  `json:"messageType"`
	Id             int    `json:"id"`
	FriendUsername string `json:"friendUsername"`
}

func (m *MessageRequest) GetNeedSearch() interface{} {
	return *m
}
