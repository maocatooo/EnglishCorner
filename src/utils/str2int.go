package utils

import "strconv"

func StrToInt(str string) (i int, err error) {
	d, err := strconv.ParseInt(str, 10, 64)
	i = int(d)
	return
}
