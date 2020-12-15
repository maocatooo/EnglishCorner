package helps

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	"EnglishCorner/spider/youdict"
	"EnglishCorner/utils/log"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var s sync.WaitGroup

//func ImportFile(path, fileName string) {
//	var (
//		DB    = db.GetDB()
//		words []string
//		w     []models.Word
//		wq    = make(map[string]string)
//	)
//	file, err := os.Open(path + fileName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	reader := bufio.NewReader(file)
//	for {
//		str, err := reader.ReadString('\n')
//		if err == io.EOF { // io.EOF 文件末尾
//			break
//		}
//		str = strings.ToLower(str)
//		str = strings.Trim(str, "\n")
//		str = strings.Trim(str, "\r")
//		str = strings.Trim(str, "\r\n")
//		words = append(words, str)
//	}
//	res := DB.Where("name IN ?", words).Find(&w)
//	var notInWords []string
//	for _, v := range w {
//		wq[v.Name] = v.Name
//	}
//	for _, k := range words {
//		_, ok := wq[k]
//		if !ok {
//			notInWords = append(notInWords, k)
//		}
//	}
//	ln := strings.TrimSuffix(fileName, ".txt")
//	l := models.CreateLibrary(ln)
//	for _, v := range w {
//		if !v.Status {
//			s.Add(1)
//			name := v.Name
//			go func() {
//				//fmt.Println(name, " --- ")
//				runSave(name, l)
//				s.Done()
//			}()
//		}
//	}
//
//	//fmt.Println(notInWords)
//	fmt.Println("length", len(words))
//	fmt.Println(res.RowsAffected)
//	//fmt.Println(res.Error)
//	s.Wait()
//}

func ImportFile(path, fileName string) {
	var (
		words []string
		DB    = db.GetDB()
	)
	file, err := os.Open(path + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 文件末尾
			break
		}
		str = strings.ToLower(str)
		str = strings.Trim(str, "\n")
		str = strings.Trim(str, "\r")
		str = strings.Trim(str, "\r\n")
		words = append(words, str)
	}
	fmt.Println(words[0])
	ln := strings.TrimSuffix(fileName, ".txt")
	l := models.CreateLibrary(ln)
	l.Public = true
	DB.Select("Name", "Public").Create(l)

	for _, v := range words {
		name := v
		s.Add(1)
		go func() {
			//fmt.Println(name, " --- ")
			runSave(name, l)
			s.Done()
		}()
	}
	s.Wait()
}

func runSave(word string, l *models.Library) {
	var (
		w  models.Word
		DB = db.GetDB()
	)
	// todo 查询数据 word 在不在数据库 循环查询数据库
	_ = DB.First(&w, "name=?", word)
	if w.Name == "" {
		log.Errorf("s:%#v, n:%#v, w:%#v \n", w.Status, w.Name, word)
	} else {
		createWord(&w, l)
	}
}

func createWord(word *models.Word, l *models.Library) {
	var (
		DB = db.GetDB()
	)
	m, _ := youdict.Find(word.Name)
	//fmt.Print(m.Status)
	if m.Status {
		w, _ := models.SetWords(*m, l, word)
		DB.Save(w)
	} else {
		fmt.Println("over! over!")
	}

}
