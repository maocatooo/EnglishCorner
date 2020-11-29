package words

import (
	"EnglishCorner/db"
	"EnglishCorner/models"
	"EnglishCorner/spider/youdict"
	"EnglishCorner/utils/resful/response"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
		"Translations").Where(
		"id >= ?", id).Limit(
		10).Find(&words, id)
	response.OK(c, words)
}

func Words(c *gin.Context) {
	var (
		words models.Word
		DB    = db.GetDB()
	)
	id, ok := c.Params.Get("id")
	if !ok {
		response.ParamsError(c, id)
		c.Abort()
	}
	res := DB.Preload(
		"Sentences").Preload(
		"Translations").First(&words, id)
	if res.Error != nil {
		response.ServerError(c, res.Error.Error())
		return
	}
	response.OK(c, words)
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
		response.ParamsError(c, err.Error())
		return
	}
	if p.Library == "" {
		response.ParamsError(c, "")
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
	response.OK(c, w)
}

func Upload(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		response.ParamsError(c, "Bad request")
		return
	}
	filename := header.Filename
	// 后缀是.txt
	ok := strings.HasSuffix(filename, ".txt")
	if !ok {
		response.ParamsError(c, "文件类型不正确")
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
	response.Tell(c, "ok")
}
