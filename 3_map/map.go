package main

import "fmt"

func main() {
	makeInit()
	secondInit()
}

//Make 方式初始化map
func makeInit() {
	m := make(map[string]int)
	m["哈哈"] = 20
	m["嘿嘿"] = 30
	fmt.Println(m)
}

func secondInit() {
	m := map[string]int{"张三":20, "嘿嘿":21}
	fmt.Println(m)
}