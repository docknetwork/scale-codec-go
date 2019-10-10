package main

//
//// #include <stdio.h>
//// #include <errno.h>
//import "C"
//import (
//	"scale/base"
//	"scale/types/compact"
//)
//
////export CDecodeUInt32
//func CDecodeUInt32(cv *C.char) (ci C.uint, cerr C.char) {
//	bytes, err := base.ConvertToBytes(C.GoString(cv))
//	if err != nil {
//		cerr = *C.CString(err.Error())
//		return
//	}
//	i, err := compact.BytesToUInt32(bytes)
//	if err != nil {
//		cerr = *C.CString(err.Error())
//		return
//	}
//	ci = C.uint(i)
//	return
//}
//
//export CDecodeUIntBig
//func CDecodeUIntBig(cv *C.char) (ci C.uint, cerr C.char) {
//	bytes, err := base.ConvertToBytes(C.GoString(cv))
//	if err != nil {
//		cerr = *C.CString(err.Error())
//		return
//	}
//	i, err := compact.BytesToUIntBig(bytes)
//	if err != nil {
//		cerr = *C.CString(err.Error())
//		return
//	}
//	ci = C.uint(i)
//	return
//}
//
//func main() {}
