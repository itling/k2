package constant

const (
	OrderByCreatedAtAsc  = "`created_at` ASC"
	OrderByCreatedAtDesc = "`created_at` DESC"
	OrderByUpdatedAtDesc = "`updated_at` DESC"
	
	//好友申请状态
	FRIEND_REQ_WAITING=1 //等待
	FRIEND_REQ_PASS=2 //通过
	FRIEND_REQ_REJECT=3 //拒绝

	HEAT_BEAT = "heatbeat"
	PONG      = "pong"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// 消息队列类型
	GO_CHANNEL = "gochannel"
	KAFKA      = "kafka"
)
