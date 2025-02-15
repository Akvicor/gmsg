package gmodel

type Status string

const (
	StatusOK     Status = "ok"
	StatusWait   Status = "wait"
	StatusCancel Status = "cancel"
	StatusFailed Status = "failed"
)

type StatusModel struct {
	Status Status `json:"status"`
	SendAt int64  `json:"send_at"`
	SentAt int64  `json:"sent_at"`
	ErrMsg string `json:"err_msg"`
}
