package main

import "fmt"

//不定参函数
func test(args ...int) {
	//fmt.Println(args)
   //len() 计算字符串有效长度
   //len() 计算不定参函数的长度
   fmt.Println(len(args))
   for i:= 0; i<len(args); i++ {
   	   fmt.Println("下标 ",i, "值 ", args[i])
   }
}

func plus(args ...int) {
	sum := 0
	/*for i:=0; i<len(args); i++ {
		sum += args[i]
	}*/

	for _,data := range args {
        sum += data
	}
	fmt.Println(sum)
}
func main()  {
   test(1,2)
   plus(1,2,3,4)


}
