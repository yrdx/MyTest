package main

import "fmt"

func main1() {
	//变量的定义与使用
	//var 变量 数据类型
	//var a int//申明
	var a int = 10
	a = a + 25
	fmt.Print(a)
}

func main() {
	//计算圆的面积和周长
	var PI float64 = 3.14
    var r float64 = 2.5
    var s float64
    var l float64
    s = PI * r * r
    l = PI *2 * r
    fmt.Println("面积", s)
    fmt.Println("周长", l)

    /*var age int
    fmt.Println("请输入你的年龄:")
    fmt.Scan(&age)
    fmt.Printf("age=%d", age)*/
    var flag bool = true
    fmt.Println(flag)
    //float32 小数点后精确7位， float64小数点后精确15位
    var f float64 = 3.14159262788
    fmt.Println(f)

    f1 := 3.14   //默认推导类型为float64
    fmt.Println(f1)

}
