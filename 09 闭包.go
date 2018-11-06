package main

import "fmt"

//函数在调用结束，会在内存中释放
func test11(a int) {
	a ++
	fmt.Println(a)
}
func main01() {
	a := 1
	for i:=0;i <10; i++ {
		test11(a)
	}
}
//可以通过匿名函数和闭包，实现函数在栈区的本地化
func test22() func() int {
    var a int
    return func() int {
    	a++
    	return a
	}
}
func main() {
	//将test22的函数类型赋值给发
  // f:= test22
   //函数调用,将test2的返回值给fn
   fn := test22()
	for i:=0;i <10; i++ {
		fmt.Print(fn())
	}

}