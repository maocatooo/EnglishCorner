package models

// 词本
type Library struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Words []*Word `json:"words" gorm:"many2many:library_words;"`
}

func (Library) TableName() string {
	return "librarys"
}

func CreateLibrary(name string) (l *Library) {
	l = new(Library)
	l.Name = name
	return
}
