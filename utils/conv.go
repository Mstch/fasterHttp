package utils

import "unsafe"

type _slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
type _string struct {
	str unsafe.Pointer
	len int
}

func ConvBytesToString(buf []byte) (str string) {
	bufp := (*_slice)(unsafe.Pointer(&buf))
	(*_string)(unsafe.Pointer(&str)).str = bufp.array
	(*_string)(unsafe.Pointer(&str)).len = bufp.len
	return
}
func ConvStringToBytes(str string) (buf []byte) {
	bufp := (*_slice)(unsafe.Pointer(&buf))
	bufp.array = (*_string)(unsafe.Pointer(&str)).str
	bufp.len = (*_string)(unsafe.Pointer(&str)).len
	bufp.cap = (*_string)(unsafe.Pointer(&str)).len
	return
}
