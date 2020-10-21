package helps

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/models"
	"bufio"
	"fmt"
	"io"
	"os"
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
	for _, v := range w {
		if !v.Status {
			s.Add(1)
			name := v.Name
			go func() {
				fmt.Printf("%s ", name)
				s.Done()
			}()
		}
	}

	//fmt.Println(notInWords)
	fmt.Println(len(words))
	fmt.Println(res.RowsAffected)
	s.Wait()
}
