package gmodel

import "encoding/json"

type RespCode int

const (
	RespCodeSucceeded       RespCode = 0 // 成功
	RespCodeFailed          RespCode = 1 // 失败
	RespCodeNotFound        RespCode = 2 // 未找到
	RespCodeBadRequest      RespCode = 3 // 错误输入
	RespCodeUnAuthorized    RespCode = 4 // 未登录
	RespCodeForbidden       RespCode = 5 // 未授权
	RespCodeConflict        RespCode = 6 // 数据冲突
	RespCodeTooManyRequests RespCode = 7 // 太多请求
)

type RespRaw struct {
	Code RespCode        `json:"code"`
	Msg  string          `json:"msg,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

type Resp struct {
	Code RespCode `json:"code"`
	Msg  string   `json:"msg,omitempty"`
	Data any      `json:"data,omitempty"`
}

func (m *Resp) DataInt64() int64 {
	switch v := m.Data.(type) {
	case int:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case uint8:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int64:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}

func (m *Resp) DataString() string {
	switch v := m.Data.(type) {
	case string:
		return v
	default:
		return ""
	}
}
