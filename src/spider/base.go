package spider

import "net/http"

type SpiderBody struct {
	KeyWord string
	LiJu    map[string]string
	CiYi    []string
	MSounds string
	YSounds string
	Status  bool
}

func AddHeader(h http.Header) {
	h["User-Agent"] = []string{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)"}
}
