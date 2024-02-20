package proto

type Msg struct {
	SenderUid    int64  `json:"sender_uid"`
	ReceiverUid  int64  `json:"receiver_uid"`
	Content      string `json:"content"`
	TimeMillUnix int64  `json:"time_mill_unix"`
	Now          string `json:"now"`
}

type Sender struct {
}

type OneChat struct {
	Chats          []*Msg  `json:"chats"`
	Sender         *Sender `json:"sender"`
	ReadAtTimeUnix int64   `json:"read_at_time_unix"`
}
