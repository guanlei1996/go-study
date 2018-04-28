package main

import (
	"fmt"
	"github.com/guanlei1996/go-study/5_type/entity"
)

func main() {
	makeType()
	valueTrans()
	pointerTransfer()
}

//创建一个简单的类
func makeType() entity.Person {
	fmt.Println("make a type")
	tom := entity.Person{Age: 10, Name: "tom"}
	fmt.Println(tom)
	return tom
}

// 类的值传递 将对象以及内部属性都拷贝一份进行传递, 在其他函数内修改不会影响到其他对象
func valueTrans() {
	fmt.Println("value transfer")
	tom := makeType()
	fmt.Println(tom)
	modifyWithValueTrans(tom)
	fmt.Println(tom)
}

func modifyWithValueTrans(person entity.Person) {
	person.Name = "Lay"
}

//类的指针传递 只传递对象的指针, 在其他函数内修改会影响到其他的引用.
func pointerTransfer() {
	fmt.Println("pointer transfer")
	tom := makeType()
	fmt.Println(tom)
	modifyWithPointerTransfer(&tom)
	fmt.Println(tom)
}

func modifyWithPointerTransfer(person *entity.Person) {
	person.Name = "Lay"
}
