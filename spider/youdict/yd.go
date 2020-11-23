package youdict

import (
	"EnglishCorner/spider"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var uri = "https://www.youdict.net/w/%s"

//var uri = "https://www.youdict.net/liju/w/%s"

func GetBody(kw string) (bs []byte, err error) {
	// 请求网页数据
	var url = fmt.Sprintf(uri, kw)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	spider.AddHeader(req.Header)

	resp, err := client.Do(req)

	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func FindLiJu(b []byte) map[string]string {
	// 查询例句
	var regx = `<div class="row" id="yd-liju">([\s\S]*?)</div>`

	reg := regexp.MustCompile(regx)
	bs := reg.Find(b)
	r := strings.NewReader(string(bs))
	d, _ := goquery.NewDocumentFromReader(r)

	var (
		liju  []string
		jiexi []string
	)
	d.Find("dt").Each(func(i int, s *goquery.Selection) {
		//解析<dt>标签
		t := s.Text()
		if t != "" && t != " " {
			liju = append(liju, t)
		}
	})
	d.Find("dd").Each(func(i int, s *goquery.Selection) {
		//解析<dd>标签
		t := s.Text()
		if t != "" && t != " " {
			jiexi = append(jiexi, t)
		}
	})
	var m = make(map[string]string)
	for i, v := range liju {
		s := strings.SplitN(v, ".", 2)
		v = strings.Trim(s[1], " ")
		if len(s) >= 2 {
			m[v] = jiexi[i]
		}
	}
	return m
}

func FindCiYi(b []byte) (y string, m string, CiYi []string) {

	defer func() {
		recover()
	}()
	r := strings.NewReader(string(b))
	d, _ := goquery.NewDocumentFromReader(r)
	d.Find("div#yd-word-pron").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		ss := strings.SplitN(strings.Trim(t, " "), "\n", 2)
		y = ss[0]
		m = ss[1]
	})
	d.Find("div#yd-word-meaning li").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		CiYi = append(CiYi, t)
	})
	return
}

func Find(kw string) (sb *spider.SpiderBody, err error) {
	sb = &spider.SpiderBody{}
	sb.KeyWord = kw
	b, e := GetBody(kw)
	if e != nil {
		return
	}
	sb.Status = true
	m := FindLiJu(b)
	sb.LiJu = m
	Y, M, C := FindCiYi(b)
	sb.CiYi = C
	sb.YSounds = Y
	sb.MSounds = M
	return
}
