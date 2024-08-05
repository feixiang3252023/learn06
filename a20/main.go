package main

import "fmt"

func main() {
	//创建一个空的map,key是int类型，value是float32类型
	map1 := make(map[int]float32)

	//向map1中添加key-value对
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	//遍历map1 , 读取key 和 value
	for key, value := range map1 {
		//打印key和 value
		fmt.Printf("key is : %d - value is : %f\n", key, value)
	}

	//遍历map1，只读取key
	for key := range map1 {
		//打印key
		fmt.Printf("key is : %d\n", key)
	}

	//遍历map1,只读取value
	for _, value := range map1 {
		//打印value
		fmt.Printf("value is : %f\n", value)
	}
}
