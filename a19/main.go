package main

import "fmt"

//声明一个包含2的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//遍历pow 切片，i是索引，v是值
	for i, v := range pow {
		//打印2的i次方等于v
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
