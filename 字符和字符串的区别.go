package main

import "fmt"

func main() {
	//var a byte = 'a'
	//var b string = "a"

	var str = "hello world"
	//计算字符串个数
	num := len(str)
	fmt.Println(num)

	//在go语言中一个汉字有3个字符
	var str1 = "中国"
	num1 := len(str1)
	fmt.Println(num1)
}
