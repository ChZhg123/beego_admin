package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x",h.Sum(nil))
}

func DateTimeToString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Date() string {
	return time.Now().Format("2006")
}

func SliceInt64ToString(arr []int64) []string {
	var sliceString []string
	for _,v := range arr{
		sliceString = append(sliceString, strconv.FormatInt(v,10))
	}
	return sliceString
}

func IsSlice(slice []string,str string) bool {
	for _,v := range slice{
		if v == str{
			return true
		}
	}
	return false
}

func Base64Encode(str string) string {
	input := []byte(str)
	return base64.StdEncoding.EncodeToString(input)
}

func Base64Decode(str string) string {
	input,_ := base64.StdEncoding.DecodeString(str)
	return string(input)
}
