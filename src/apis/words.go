package apis

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/models"
	"EnglishCorner/src/spider/youdict"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func WordsList(c *gin.Context) {
	var (
		words []models.Word
		DB    = db.GetDB()
	)
	id, ok := c.GetQuery("id")
	if !ok {
		id = "1"
	}
	DB.Preload(
		"Sentences").Preload(
		"Translations").Preload(
		"Translations").Where(
		"id >= ?", id).Limit(
		10).Find(&words, id)
	c.JSON(200, words)
}

func Words(c *gin.Context) {
	var (
		words models.Word
		DB    = db.GetDB()
	)
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(422, words)
	}
	res := DB.First(&words, id)
	if res.Error != nil {
		c.JSON(422, gin.H{})
		return
	}
	c.JSON(200, words)
}

type postWords struct {
	Name    string `json:"name"`
	Library string `json:"library"`
}

func WordsPost(c *gin.Context) {
	var (
		p  postWords
		DB = db.GetDB()
	)
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(422, gin.H{})
		return
	}
	if p.Library == "" {
		c.JSON(422, gin.H{})
		return
	}
	//fmt.Println(c.Keys)
	m, _ := youdict.Find(p.Name)
	var l models.Library

	res := DB.Where("name = ?", p.Library).First(&l)

	if res.RowsAffected < 1 {
		l.Name = p.Library
		DB.Create(&l)
	}
	w, _ := models.CreateWords(*m, &l)
	DB.Create(w)
	c.JSON(200, w)
}

func Upload(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename
	// 后缀是.txt
	ok := strings.HasSuffix(filename, ".txt")
	if !ok {
		c.String(http.StatusBadRequest, "文件类型不正确")
		return
	}
	// 保存文件到本地
	//out, err := os.Create(filename)
	//if err != nil {
	//	log.Error(err)
	//}
	//defer out.Close()
	//_, err = io.Copy(out, file)
	//if err != nil {
	//	log.Error(err)
	//}
	reader := bufio.NewReader(file)
	var words []string
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF { // io.EOF 文件末尾
			break
		}
		words = append(words, str)
	}
	c.JSON(http.StatusOK, gin.H{})
}
