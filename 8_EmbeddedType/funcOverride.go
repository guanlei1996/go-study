package main

import "github.com/guanlei1996/go-study/8_EmbeddedType/entity"

func main() {
	ad := entity.Admin{User: entity.User{Name: "lay"}, Level: "administrator"}
	ad.User.SayHello()
	// 这次执行的方法为覆盖后的方法
	ad.SayHello()
}
