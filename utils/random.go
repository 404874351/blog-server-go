package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//
// RandomNumbers
//  @Description: 生成指定长度的随机数字序列
//  @param length
//  @return string
//
func RandomNumbers(length int) string {
	if length < 1 {
		length = 1
	}
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}