package gmsg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akvicor/gmsg/gmodel"
)

func Cancel(ctx context.Context, id int64) error {
	body := gmodel.NewSendCancel(id)
	header := map[string]string{
		accessToken: AccessToken.String(),
	}
	data, err := cli.Post(ctx, ApiUrl.Append("/api/send/cancel"), body, header)
	if err != nil {
		return fmt.Errorf("send error: %v", err)
	}
	resp := &gmodel.Resp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("send error: %v", err)
	}
	if resp.Code == gmodel.RespCodeSucceeded {
		return nil
	}
	return fmt.Errorf("cancel error code %d, %s", resp.Code, resp.Msg)
}
