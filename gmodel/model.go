package gmodel

type Send struct {
	ID    int64  `json:"id" form:"id" query:"id"`
	Sign  string `json:"sign" form:"sign" query:"sign"` // 唯一标记
	Type  Type   `json:"type" form:"type" query:"type"` // 消息类型
	At    int64  `json:"at" form:"at" query:"at"`       // SendAt
	Title string `json:"title" form:"title" query:"title"`
	Msg   string `json:"msg" form:"msg" query:"msg"`
}

type SendCancel struct {
	ID int64 `json:"id" form:"id" query:"id"`
}

type SendStatus struct {
	ID int64 `json:"id" form:"id" query:"id"`
}

func NewSend(id int64, sign string, cType Type, at int64, title, msg string) *Send {
	return &Send{
		ID:    id,
		Sign:  sign,
		Type:  cType,
		At:    at,
		Title: title,
		Msg:   msg,
	}
}

func NewSendCancel(id int64) *SendCancel {
	return &SendCancel{
		ID: id,
	}
}

func NewSendStatus(id int64) *SendStatus {
	return &SendStatus{
		ID: id,
	}
}
