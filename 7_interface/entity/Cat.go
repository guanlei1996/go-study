package entity

import "fmt"

type Cat struct {

}

func (cat Cat) SayHello() {
	fmt.Println("喵喵喵~")
}

func (cat Cat) Run() {
	fmt.Println("喵~")
}