package swos

import (
	"encoding/hex"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func unmarshalBody(body *[]byte, v interface{}) error {
	vReflect := reflect.TypeOf(v)
	for i := 0; i < vReflect.Elem().NumField(); i++ {
		rType := vReflect.Elem().Field(i)
		rVal := reflect.ValueOf(v).Elem().Field(i)

		if rType.Type.Kind() != reflect.Slice {
			switch rType.Type.String() {
			case "uint8":
				val := findstring(body, strings.ToLower(rType.Name))
				rVal.SetUint(tpUint(val, 8))
			case "uint32":
				val := findstring(body, strings.ToLower(rType.Name))
				rVal.SetUint(tpUint(val, 32))
			case "string":
				val := findstring(body, strings.ToLower(rType.Name))
				rVal.SetString(val)
			default:

			}
		}
		if rType.Type.Kind() == reflect.Slice {

			switch rType.Type.String() {
			case "[]uint8":

				val := findstring(body, strings.ToLower(rType.Name))
				var slc []uint8
				for _, valSi := range strings.Split(val, ",") {
					slc = append(slc, uint8(tpUint(valSi, 8)))
				}
				refslice := reflect.MakeSlice(reflect.TypeOf(slc), len(slc), cap(slc))
				for si, sv := range slc {
					refslice.Index(si).SetUint(uint64(sv))
				}
				rVal.Set(refslice)
			case "[]uint16":

				val := findstring(body, strings.ToLower(rType.Name))
				var slc []uint16
				for _, valSi := range strings.Split(val, ",") {
					slc = append(slc, uint16(tpUint(valSi, 16)))
				}
				refslice := reflect.MakeSlice(reflect.TypeOf(slc), len(slc), cap(slc))
				for si, sv := range slc {
					refslice.Index(si).SetUint(uint64(sv))
				}
				rVal.Set(refslice)
			case "[]string":
				val := findstring(body, strings.ToLower(rType.Name))
				var slc []string
				for _, valSi := range strings.Split(val, ",") {
					slc = append(slc, getStringByteString(valSi))
				}
				refslice := reflect.MakeSlice(reflect.TypeOf(slc), len(slc), cap(slc))
				for si, sv := range slc {
					refslice.Index(si).SetString(sv)
				}
				rVal.Set(refslice)
			default:
				log.Println(rType.Type)
			}
		}

	}

	return nil
}

func getStringByteString(s string) string {
	if s[:1] == "'" {
		s = s[1:]
	}
	if s[len(s)-1:] == "'" {
		s = s[:len(s)-1]
	}
	vl, _ := hex.DecodeString(s)
	return string(vl)
}

func tpUint(s string, bitsize int) uint64 {
	clear := strings.Replace(s, "0x", "", -1)
	result, _ := strconv.ParseUint(clear, 16, bitsize)
	return result
}

func findstring(b *[]byte, key string) string {
	value := findvaluestring(b, key)
	return value
}

func findvaluestring(b *[]byte, k string) string {
	str := string(*b)
	key := k + ":"
	startIndex := strings.Index(str, key)
	if startIndex == -1 {
		return ""
	}
	keyLen := len(key)
	endIndex := startIndex + keyLen
	firstcharsetvalue := str[endIndex : endIndex+1]
	if firstcharsetvalue == "[" {
		vlIStart := endIndex + 1
		vlIEnd := strings.Index(str[vlIStart:], "]")
		if vlIEnd == -1 {
			return ""
		}
		return str[vlIStart : vlIStart+vlIEnd]

	}
	if firstcharsetvalue == "'" {
		vl_i_start := endIndex + 2
		vl_i_end := strings.Index(str[vl_i_start:], "'")
		if vl_i_end == -1 {
			return ""
		}
		return str[vl_i_start:vl_i_end]

	}
	nextZap := strings.Index(str[endIndex:], ",")
	return str[endIndex : endIndex+nextZap]

}

func getBoolFromUint32(u uint32, i uint8) bool {
	a := u >> i
	b := a & 0x1
	if b == 0 {
		return false
	}
	return true
}

func getLacp(u uint8, i uint8) {

}
