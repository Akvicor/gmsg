# GMSG

[MSG](https://github.com/Akvicor/msg) 的Go Mod

## 配置

```go
// 访问密钥
AccessToken
// API地址
ApiUrl

// 修改方式
gmsg.AccessToken.Set()
```

## 参数

```go
// 普通文本
TypeText     = model.TypeText
// 微信文本卡片
TypeTextCard = model.TypeTextCard
// Markdown
TypeMarkdown = model.TypeMarkdown
// HTML
TypeHTML     = model.TypeHTML
```

## 方法

```go
// 取消还未发送的消息
Cancel(ctx context.Context, id int64) error
// 通过渠道ID发送
SendByID(ctx context.Context, id int64, cType model.Type, at int64, title, msg string) (int64, error)
// 通过渠道标记发送
SendBySign(ctx context.Context, sign string, cType model.Type, at int64, title, msg string) (int64, error)
// 通过消息ID获取发送状态
Status(ctx context.Context, id int64) (*model.StatusModel, error)
```
