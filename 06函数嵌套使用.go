package main

import "fmt"

func test1(a int, b int) {
   test2(a, b)
}

func test2(a int, b int) {
    sum := a + b
    fmt.Println(sum)
}
//所有的函数都是全局函数，可以被项目中所有文件使用 在项目中函数名是唯一的
func main()  {
	test1(10,20)
}
