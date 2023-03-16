/**
  @author: Zero
  @date: 2023/3/16 20:38:37
  @desc: 随机生成工具

**/

package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const NUMBERS = "0123456789"

// RandomMobileCode 生成一个6位数手机验证码
func RandomMobileCode() string {
	return RandomNumberCode(6)
}

// RandomNumberCode 随机生成指定长度的数字字符
func RandomNumberCode(length int) string {
	numLength := len(NUMBERS)
	// 设置一个随机种子
	rand.Seed(time.Now().UnixNano())
	var code bytes.Buffer
	//var code strings.Builder
	for i := 0; i < length; i++ {
		// 每次随机从0-9获取一个字符,追加到code
		code.WriteByte(NUMBERS[rand.Intn(numLength)])
	}
	return code.String()
}

func main() {
	fmt.Println(RandomNumberCode(6))
}
