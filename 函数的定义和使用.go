package main

import "fmt"

//函数定义
//func 函数名(函数参数列表) {
//   代码体 程序体 函数体
// }
//变量名 数据类型
//函数定义
func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func swap(a int, b int) {
	a, b = b, a
}
func main() {
   //函数调用
   add(1, 2)


}
