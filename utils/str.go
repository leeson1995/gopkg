/*******************************************************************************
 * // Author: leeson9616@gmail.com
 * // Update：
 ******************************************************************************/

package utils

import (
	"encoding/json"
	"golang.org/x/exp/rand"
	"reflect"
	"strconv"
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

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
func StringToInt64(i string) int64 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return j
}

//judge a string whether in the  string list
func IsContain(target string, List []string) bool {
	for _, element := range List {

		if target == element {
			return true
		}
	}
	return false

}

func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}
func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

//The incoming parameter must be a pointer
func JsonStringToStruct(s string, args interface{}) error {
	err := json.Unmarshal([]byte(s), args)
	return err
}

func GetMsgID(sendID string) string {
	t := int64ToString(GetCurrentTimestampByNano())
	return Md5(t + sendID + int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}
func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
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
