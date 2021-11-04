package swos

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func marshalToBody(v interface{}) (dst []byte, err error) {
	dstVal := reflect.ValueOf(v)
	dstType := reflect.TypeOf(v)

	var res string
	for i := 0; i < dstType.NumField(); i++ {
		rType := dstType.Field(i)
		rVal := dstVal.Field(i)

		res += strings.ToLower(rType.Name) + ":"
		if rType.Type.Kind() != reflect.Slice {
			switch rType.Type.String() {
			case "uint8":
				res += fmt.Sprintf("0x%02s", strconv.FormatUint(rVal.Uint(), 16))
			case "uint16":
				res += fmt.Sprintf("0x%04s", strconv.FormatUint(rVal.Uint(), 16))
			case "uint32":
				res += fmt.Sprintf("0x%08s", strconv.FormatUint(rVal.Uint(), 16))
				//res += "0x" + strconv.FormatUint(rVal.Uint(), 16)
			case "string":
				res += encodeString(rVal.String())
			default:
			}
		} else {
			res += "["
			switch rType.Type.String() {
			case "[]uint8":
				sl := rVal.Interface().([]uint8)
				for _, v := range sl {
					res += fmt.Sprintf("0x%02s,", strconv.FormatUint(uint64(v), 16))
					//res += "0x" + fmt.Sprintf("%02.0f",float32(v)) + ","
					//res += "0x" + strconv.FormatUint(uint64(v), 16) + ","
				}
			case "[]uint16":
				sl := rVal.Interface().([]uint16)
				for _, v := range sl {
					res += fmt.Sprintf("0x%04s,", strconv.FormatUint(uint64(v), 16))
					//res += "0x" + strconv.FormatUint(uint64(v), 16) + ","
				}
			case "[]uint32":
				sl := rVal.Interface().([]uint32)
				for _, v := range sl {
					res += fmt.Sprintf("0x%08s,", strconv.FormatUint(uint64(v), 16))
					//res += "0x" + strconv.FormatUint(uint64(v), 16) + ","
				}
			case "[]string":
				sl := rVal.Interface().([]string)
				for _, v := range sl {
					res += encodeString(v) + ","
				}
			default:
			}
			res = res[0:len(res)-1] + "]"
		}
		res += ","

	}
	dst = []byte("{" + res[0:len(res)-1] + "}")
	return dst, nil
}
func encodeString(src string) string {
	return fmt.Sprintf("'%s'", hex.EncodeToString([]byte(src)))
}
