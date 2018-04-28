package main

import "fmt"

func main() {
	ad:= admin{user{"lay", "guanlei@geeweb.cn"}, "1"}
	fmt.Println(ad.name)
	fmt.Println(ad.email)
	fmt.Println(ad.level)
	ad.SayHello()
}

type user struct {
	name string
	email string
}

func (u user) SayHello()  {
	fmt.Printf("Hello, I'm %s. ", u.name)
}

type admin struct {
	//user就是admin的嵌入类型, admin拥有user的所有属性字段/方法等等所有.
	user
	level string
}