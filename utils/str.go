/*******************************************************************************
 * // Author: leeson9616@gmail.com
 * // Update：
 ******************************************************************************/

package utils

import (
	"reflect"
	"unsafe"
)

func String2Bytes(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func Bytes2String(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}


//FYR gin
//func StringToBytes(s string) []byte {
// return *(*[]byte)(unsafe.Pointer(
//  &struct {
//   string
//   Cap int
//  }{s, len(s)},
// ))
//}
//
//func BytesToString(b []byte) string {
// return *(*string)(unsafe.Pointer(&b))
//}