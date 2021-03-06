package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// StringToInt64 字符串转换成int64
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToInt 字符串转int
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// GetCurrentTimeStr
// 获取当前时间字符串
func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// StructToJsonStr 结构体转json字符串
func StructToJsonStr(e interface{}) (string, error) {
	if marshal, err := json.Marshal(e); err == nil {
		return string(marshal), err
	} else {
		return "", err
	}
}

// JsonStrToMap json 字符串转map
func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

// StructToMap struct 转 map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

// GetValidateCode 根据长度获取随机字符串
func GetValidateCode(width int) string {
	numeric := []byte("0123456789abcdefghijklmnopqrstuvwzyxABCDEFGHIJKLMNOPQRESTUVWXYZ")
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// GetRandString 获取指定长度随机字符串
func GetRandString(length int64) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	runes := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range letters {
		runes[i] = letters[rand.Intn(len(letters))]
		if i == int(length)-1 {
			break
		}
	}
	return string(runes)
}

// GetRandStringByRune 指定字符串种子获取随机字符串
func GetRandStringByRune(length int, letters []rune) string {
	runes := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range letters {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}

func GetEncryptStringByMd5(passwordStr string, encryptStr string) string {
	bytes := []byte(passwordStr + encryptStr)
	hash := md5.New()
	hash.Write(bytes)

	return hex.EncodeToString(hash.Sum(nil))
}
