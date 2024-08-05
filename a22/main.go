package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	//忽略索引
	for _, num := range nums {
		fmt.Println("value:", num)
	}
	//忽略值
	for i := range nums {
		fmt.Println("index:", i)
	}
}
