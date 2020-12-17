package models

// 词本
type Library struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Words  []*Word `json:"words" gorm:"many2many:library_words;"`
	UserID int     `json:"user_id"`
	Public bool    `json:"public"`
}

func (Library) TableName() string {
	return "librarys"
}

type LibraryWords struct {
	LibraryID uint `json:"library_id"`
	WordID    uint `json:"word_id"`
	LWID      uint `json:"lwid" gorm:"unique;autoIncrement:true"`
}

func (LibraryWords) TableName() string {
	return "library_words"
}

func CreateLibrary(name string) (l *Library) {
	l = new(Library)
	l.Name = name
	return
}
