package pkg

import (
	"EnglishCorner/utils"
	"bytes"
	"encoding/base64"
	"github.com/dchest/captcha"
	"time"
)

func GetKey() string {

	t := time.Now().UnixNano()
	return utils.Int64ToStr(t)

}

func Captcha(length ...int) (string, string) {
	l := captcha.DefaultLen - 2
	w, h := 200, 50
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	//_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
	image := Serve(captchaId, w, h)
	return captchaId, image
}

//
//func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
//	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
//	w.Header().Set("Pragma", "no-cache")
//	w.Header().Set("Expires", "0")
//
//
//	var content bytes.Buffer
//	switch ext {
//	case ".png":
//		w.Header().Set("Content-Type", "image/png")
//		_ = captcha.WriteImage(&content, id, width, height)
//	case ".wav":
//		w.Header().Set("Content-Type", "audio/x-wav")
//		_ = captcha.WriteAudio(&content, id, lang)
//	default:
//		return captcha.ErrNotFound
//	}
//	if download {
//		w.Header().Set("Content-Type", "application/octet-stream")
//	}
//	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
//	return nil
//}

func Serve(id string, width, height int) string {
	var content bytes.Buffer
	captcha.WriteImage(&content, id, width, height)

	return base64.StdEncoding.EncodeToString(content.Bytes())
}

func CaptchaVerify(captchaId, code string) bool {
	return captcha.VerifyString(captchaId, code)
}
