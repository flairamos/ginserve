package random_util

import (
	"fmt"
	"math/rand"
)

// 随机生成4个数字
func RandInt4i() int {
	base := 8
	first := (rand.Intn(base) + 1) * 1000
	second := (rand.Intn(base) + 1) * 100
	third := (rand.Intn(base) + 1) * 10
	return first + second + third + rand.Intn(base)
}

func RandInt3i() int {
	base := 8
	first := (rand.Intn(base) + 1) * 100
	second := (rand.Intn(base) + 1) * 10
	fmt.Println(first, second)
	return first + second + rand.Intn(base)
}

func RandInt4s() string {
	base := 10
	first := rand.Intn(base)
	second := rand.Intn(base)
	third := rand.Intn(base)
	fourth := rand.Intn(base)
	return fmt.Sprintf("%d%d%d%d", first, second, third, fourth)
}
