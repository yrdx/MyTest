package main

import "fmt"

func main() {
	a := 10
	b := 20
	//匿名内部函数
	/*func (a int, b int) {
		fmt.Println(a + b)
	}(a, b)*/

	f := func (a int , b int) {
		fmt.Println(a + b)
	}
	f(a , b)

	fn := func () {
		fmt.Print("Hello World!")
	}
	fn()
	fn()
}
