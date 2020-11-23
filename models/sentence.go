package models

import (
	"EnglishCorner/spider"
)

// 例句
type Sentence struct {
	ID     uint `json:"id"`
	WordID uint `json:"word_id"`
	// 例句
	Sent string `json:"sent"`
	// 翻译
	Trans string `json:"trans"`
}

func (Sentence) TableName() string {
	return "sentences"
}

func GetSentencesFormSpiderBody(sb spider.SpiderBody) ([]Sentence, bool) {
	if !sb.Status {
		return nil, sb.Status
	}
	var (
		s []Sentence
	)
	for k, v := range sb.LiJu {
		s = append(s, Sentence{
			Sent:  k,
			Trans: v,
		})
	}
	return s, sb.Status
}
