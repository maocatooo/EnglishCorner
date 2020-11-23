package models

import (
	"EnglishCorner/spider"
)

type Translation struct {
	ID     uint   `json:"id"`
	WordID uint   `json:"word_id"`
	Name   string `json:"name"`
}

func (Translation) TableName() string {
	return "translations"
}

func GetTranslationsFormSpiderBody(sb spider.SpiderBody) ([]Translation, bool) {
	if !sb.Status {
		return nil, sb.Status
	}
	var (
		t []Translation
	)
	for _, v := range sb.CiYi {
		t = append(t, Translation{
			Name: v,
		})
	}
	return t, sb.Status
}
