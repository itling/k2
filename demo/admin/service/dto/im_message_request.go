package dto

type ImMessageRequest struct {
	MessageType    int32  `form:"messageType"`
	Id             int    `form:"id"` //用户id或群组id
	FriendUsername string `form:"friendUsername"`
}

func (m *ImMessageRequest) GetNeedSearch() interface{} {
	return *m
}
