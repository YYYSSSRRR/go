package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCode() string {
	//rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	code := random.Intn(900000) + 100000
	//Sprintf是用来转换成字符串的，不在控制台打印
	return fmt.Sprintf("%d", code)
}
