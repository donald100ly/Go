package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt"
	"math/rand"
	"time"
)

func main() { // main函数，是程序执行的入口
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num) //     在终端打印 Hello World!
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("End")
}
