package helps

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/models"
	"EnglishCorner/src/spider/youdict"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var s sync.WaitGroup

func ImportFile(path, fileName string) {
	var (
		DB    = db.GetDB()
		words []string
		w     []models.Word
		wq    = make(map[string]string)
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
	res := DB.Where("name IN ?", words).Find(&w)
	var notInWords []string
	for _, v := range w {
		wq[v.Name] = v.Name
	}
	for _, k := range words {
		_, ok := wq[k]
		if !ok {
			notInWords = append(notInWords, k)
		}
	}
	ln := strings.TrimSuffix(fileName, ".txt")
	l := models.CreateLibrary(ln)
	for _, v := range w {
		if !v.Status {
			s.Add(1)
			name := v.Name
			go func() {
				//fmt.Println(name, " --- ")
				runSave(name, l)
				s.Done()
			}()
		}
	}

	//fmt.Println(notInWords)
	fmt.Println("length", len(words))
	fmt.Println(res.RowsAffected)
	//fmt.Println(res.Error)
	s.Wait()
}

func runSave(word string, l *models.Library) {
	var (
		w  models.Word
		DB = db.GetDB()
	)

	_ = DB.First(&w, "name=?", word)

	if w.Name == "" {
		fmt.Printf("s:%#v, n:%#v, w:%#v \n", w.Status, w.Name, word)
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
	}

}
