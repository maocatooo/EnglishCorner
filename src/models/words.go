package models

import (
	"EnglishCorner/src/spider"
)

// 单词
type Word struct {
	ID uint `json:"id"`

	Name         string        `json:"name"`
	MSounds      string        `json:"m_sounds"`
	YSounds      string        `json:"y_sounds"`
	Status       bool          `json:"status"`
	Sentences    []Sentence    `json:"sentences"`
	Translations []Translation `json:"translations"`
	Library      []*Library    `json:"library" gorm:"many2many:library_words;"`
}

func (Word) TableName() string {
	return "words"
}

func GetWordsFormSpiderBody(sb spider.SpiderBody) (*Word, bool) {
	if !sb.Status {
		return nil, sb.Status
	}
	var (
		w Word
	)
	w.MSounds = sb.MSounds
	w.YSounds = sb.YSounds
	w.Name = sb.KeyWord
	w.Status = sb.Status
	return &w, sb.Status
}

func CreateWords(sb spider.SpiderBody, l *Library) (*Word, bool) {
	var ok bool
	words, ok := GetWordsFormSpiderBody(sb)
	if !ok {
		return nil, ok
	}
	words.Library = []*Library{
		l,
	}
	s, ok := GetSentencesFormSpiderBody(sb)
	if !ok {
		return nil, ok
	}
	words.Sentences = s
	t, ok := GetTranslationsFormSpiderBody(sb)
	words.Translations = t
	//db.Create()
	return words, ok
}
