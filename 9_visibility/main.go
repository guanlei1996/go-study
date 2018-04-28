package main

import "github.com/guanlei1996/go-study/9_visibility/entity"

func main() {
	person := entity.Person{Name: "Lay"}
	person.SayHello()
	// 下方方式不可访问
	// person.shortName
	// person.hello()
}
