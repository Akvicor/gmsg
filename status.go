package gmsg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akvicor/gmsg/gmodel"
)

func Status(ctx context.Context, id int64) (*gmodel.StatusModel, error) {
	body := gmodel.NewSendStatus(id)
	header := map[string]string{
		accessToken: AccessToken.String(),
	}
	data, err := cli.Post(ctx, ApiUrl.Append("/api/send/status"), body, header)
	if err != nil {
		return nil, fmt.Errorf("send status error: %v", err)
	}
	resp := &gmodel.RespRaw{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, fmt.Errorf("send status resp error: %v", err)
	}
	if resp.Code != gmodel.RespCodeSucceeded {
		return nil, fmt.Errorf("status error code %d, %s", resp.Code, resp.Msg)
	}
	status := &gmodel.StatusModel{}
	err = json.Unmarshal(resp.Data, status)
	if err != nil {
		return nil, fmt.Errorf("send status data error: %v", err)
	}
	return status, nil
}
