package strUtil

import (
	"reflect"
	"unsafe"
)

// return immutable []byte
func StringToSliceByte(s string) []byte {
	l := len(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  l,
		Cap:  l,
	}))
}
