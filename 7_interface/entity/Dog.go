package entity

import "fmt"

type Dog struct {

}

func (dog Dog) SayHello() {
	fmt.Println("汪汪汪~")
}

func (dog Dog) Run()  {
	fmt.Println("汪~")
}