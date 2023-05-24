package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// 定义命令行参数
	input := flag.String("input", "", "The input string")

	// 解析命令行参数
	flag.Parse()

	// 反转输入字符串
	output := reverseString(*input)

	// 输出结果
	fmt.Println("Reversed string:", output)
}

func reverseString(s string) string {
	// 将字符串转换为字符数组
	chars := strings.Split(s, "")

	// 反转字符数组
	for i := 0; i < len(chars)/2; i++ {
		j := len(chars) - 1 - i
		chars[i], chars[j] = chars[j], chars[i]
	}

	// 将字符数组转换为字符串
	return strings.Join(chars, "")
}
