package main

import "fmt"

//func 函数名(形参列表) 返回值类型列表{代码体}
func sub(a int, b int) int {
  sum := a - b
  return sum
}
//函数多个返回值
func test5() (a int, b int, c int) {
	a,b,c = 1,2,3
	return
}
func main() {
   value := sub(10, 20)
   fmt.Println(value)

   a,b,c := test5()
   fmt.Println(a,b,c)
}
