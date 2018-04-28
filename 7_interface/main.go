package main

import "github.com/guanlei1996/go-study/7_interface/entity"

func main() {
	// 声明接口类
	var a entity.Animal
	var b entity.Animal
	// 声明具体实现 并将实现赋值给接口
	var cat entity.Cat
	a = cat

	var dog entity.Dog
	b = dog

	// 执行接口方法
	a.SayHello()
	b.SayHello()
	a.Run()
	b.Run()
}
