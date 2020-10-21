package helps

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/models"
	"EnglishCorner/src/utils/log"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func CreateTable() {
	var (
		DB = db.GetDB()
	)
	_ = DB.AutoMigrate(&models.Word{})
	_ = DB.AutoMigrate(&models.Library{})
	_ = DB.AutoMigrate(&models.Sentence{})
	_ = DB.AutoMigrate(&models.Translation{})
}

func DeleteTableData() {
	var (
		DB = db.GetDB()
	)
	//DROP TABLE table_name
	//DELETE FROM table_name
	DB.Exec("DROP TABLE library_words")
	DB.Exec("DROP TABLE librarys")
	DB.Exec("DROP TABLE sentences")
	DB.Exec("DROP TABLE translations")
	DB.Exec("DROP TABLE words")
}

func InitData(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var (
		i     = 0
		words []models.Word
		DB    = db.GetDB()
	)
	// todo 去重复
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 文件末尾
			break
		}
		i += 1
		//fmt.Println(i)
		if i%2 == 0 {
			continue
		}
		words = append(words, models.Word{Name: str})

	}
	length := len(words)
	if length <= 0 {
		return
	}

	var f float64
	flag := 2000
	f = float64(length) / float64(flag)
	c := int(math.Ceil(f))
	fmt.Println(c)
	for i := 0; i < c; i++ {
		if (i + 1) == c {
			// 最后一次
			_w := words[i*flag:]
			DB.Create(&_w)
		} else {
			_w := words[i*flag : (i+1)*flag]
			DB.Create(&_w)
		}
	}
	fmt.Println(length)
	res := DB.Find(&[]models.Word{})
	fmt.Println(res.RowsAffected)
}
