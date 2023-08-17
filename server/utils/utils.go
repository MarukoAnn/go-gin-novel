package utils

import (
	"reflect"
	"time"
)

func Map(slice interface{}, fn func(interface{}, int) interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	result := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len(), sliceValue.Cap())
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		//result.Index(i).Set(reflect.ValueOf())
		fn(item.Interface(), i)
	}
	return result.Interface()
}

func TimerTask(interval int64, callback func()) {
	for {
		<-time.After(time.Second * time.Duration(interval))
		// 在这里执行定时任务的操作
		callback()
	}
}
