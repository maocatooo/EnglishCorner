package utils

import "strconv"

func IntToStr(i int) string {

	return Int64ToStr(int64(i))
}

func Int64ToStr(i int64) string {

	return strconv.FormatInt(i, 10)
}

func Int32ToStr(i int32) string {
	return Int64ToStr(int64(i))
}
