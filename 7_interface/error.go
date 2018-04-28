package main

import "fmt"

func main() {
	var a animal
	var b animal
	var c animal
	var d animal

	var valueCat cat
	var pointerCat *cat
	// 此处会出现编译错误, 因为cat实现的方法接收者为指针类型, 所以只能实现接收者为指针类型的接口
	a=valueCat
	b=pointerCat


	var valueDog dog
	var pointerDog dog
	// 此处编译可以顺利通过, 因为dog实现的方法接收者为值类型, 说明值类型既可以实现值类型接口也能实现指针类型接口
	c=valueDog
	d=pointerDog

	invoke(a)
	invoke(b)
	invoke(c)
	invoke(d)
}

func invoke(a animal) {
	a.sayHello()
}

func invokeWithPointer(a *animal) {

}
//接口类型
type animal interface {
	sayHello()
}
//实现类
type cat struct {
}
type dog struct {
}
//cat的接收者为指针类型
func (c *cat) sayHello() {
	fmt.Println("喵")
}
//dog的接收者为值类型
func (d dog) sayHello() {
	fmt.Println("汪")
}