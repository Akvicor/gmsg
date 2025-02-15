package gmsg

import "github.com/Akvicor/gmsg/gmodel"

const (
	TypeText     = gmodel.TypeText     // 普通文本
	TypeTextCard = gmodel.TypeTextCard // 微信文本卡片
	TypeMarkdown = gmodel.TypeMarkdown // Markdown
	TypeHTML     = gmodel.TypeHTML     // HTML
)

type TypeDetail struct {
	Type        gmodel.Type `json:"type"`
	Name        string      `json:"name"`
	EnglishName string      `json:"english_name"`
}

var AllType = []*TypeDetail{
	{TypeText, TypeText.String(), TypeText.StringEnglish()},
	{TypeTextCard, TypeTextCard.String(), TypeTextCard.StringEnglish()},
	{TypeMarkdown, TypeMarkdown.String(), TypeMarkdown.StringEnglish()},
	{TypeHTML, TypeHTML.String(), TypeHTML.StringEnglish()},
}
