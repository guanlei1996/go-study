package main

import "fmt"

func main() {
	testCap()
}

/**
1.经测试, 如果make的大小填入数字, 则append函数会从指定长度后开始追加
2.slice扩容的方式为 两倍递增 如 1->2  2->4  4->8 8->16...
 */
func testCap() {
	slice := make([]int, 0)

	for i := 0; i < 10; i++ {
		slice = append(slice, 1)
		fmt.Printf("index : %d\n", i)
		fmt.Printf("len : %d, cat : %d", len(slice), cap(slice))
	}
	fmt.Println(slice)
}
