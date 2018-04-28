package main

import "fmt"

func main() {
	//forRange()
	//arrayToFunc()
}

// for-range方式遍历数组
func forRange() {
	array := [5]int{1: 3, 4: 1}
	for index, value := range array {
		fmt.Printf("array index : %d\n", index)
		fmt.Printf("array value : %d\n", value)
	}
}

// 数组指针传递
func arrayToFunc() {
	array := [5]int{1, 4, 5, 1, 12}
	fmt.Println(array)
	modify(&array)
	fmt.Println(array)
}

func modify(a *[5]int) {
	a[0] = 3
	fmt.Println(*a)
}
