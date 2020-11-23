package utils

import (
	"math/rand"
	"reflect"
	"time"
)

// 打乱切片的顺序
func RandSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	length := rv.Len()
	if length < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := length/2 + 1; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}
