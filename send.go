package gmsg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akvicor/gmsg/gmodel"
)

func SendByID(ctx context.Context, id int64, cType gmodel.Type, at int64, title, msg string) (int64, error) {
	body := gmodel.NewSend(id, "", cType, at, title, msg)
	header := map[string]string{
		accessToken: AccessToken.String(),
	}
	data, err := cli.Post(ctx, ApiUrl.Append("/api/send"), body, header)
	if err != nil {
		return 0, fmt.Errorf("send error: %v", err)
	}
	resp := &gmodel.Resp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return 0, fmt.Errorf("send error: %v", err)
	}
	if resp.Code == gmodel.RespCodeSucceeded {
		msgId := resp.DataInt64()
		if msgId == 0 {
			return 0, errors.New("wrong id")
		}
		return msgId, nil
	}
	return 0, fmt.Errorf("send error code %d:%s", resp.Code, resp.Msg)
}

func SendBySign(ctx context.Context, sign string, cType gmodel.Type, at int64, title, msg string) (int64, error) {
	body := gmodel.NewSend(0, sign, cType, at, title, msg)
	header := map[string]string{
		accessToken: AccessToken.String(),
	}
	data, err := cli.Post(ctx, ApiUrl.Append("/api/send"), body, header)
	if err != nil {
		return 0, fmt.Errorf("send error: %v", err)
	}
	resp := &gmodel.Resp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return 0, fmt.Errorf("send error: %v", err)
	}
	if resp.Code == gmodel.RespCodeSucceeded {
		msgId := resp.DataInt64()
		if msgId == 0 {
			return 0, errors.New("wrong id")
		}
		return msgId, nil
	}
	return 0, fmt.Errorf("send error code %d:%s", resp.Code, resp.Msg)
}
