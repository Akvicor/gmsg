package gmsg

import (
	"context"
	"github.com/Akvicor/gmsg/gmodel"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	id, err := SendBySign(context.Background(), "telegram", gmodel.TypeMarkdown, time.Now().Unix()+3, "title", "content "+time.Now().String())
	if err != nil {
		t.Error(err)
	}
	t.Logf("msg id: %d", id)
	status, err := Status(context.Background(), id)
	if err != nil {
		t.Error(err)
	}
	if status.SentAt < 0 {
		status.SentAt = -status.SentAt
	}
	t.Logf("status: [%v], sendAt: [%s], sentAt: [%s], errMsg: [%s]", status.Status, time.Unix(status.SendAt, 0).Format(time.DateTime), time.Unix(status.SentAt, 0).Format(time.DateTime), status.ErrMsg)
	err = Cancel(context.Background(), id)
	if err != nil {
		t.Error(err)
	}
	t.Log("cancel ok")
	status, err = Status(context.Background(), id)
	if err != nil {
		t.Error(err)
	}
	if status.SentAt < 0 {
		status.SentAt = -status.SentAt
	}
	t.Logf("status: [%v], sendAt: [%s], sentAt: [%s], errMsg: [%s]", status.Status, time.Unix(status.SendAt, 0).Format(time.DateTime), time.Unix(status.SentAt, 0).Format(time.DateTime), status.ErrMsg)
}
