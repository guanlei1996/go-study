package entity

import (
	"fmt"
)

type Person struct {
	Name string
	shortName string
}

func (p Person) SayHello() {
	p.hello()
	p.shortName = p.Name[0:len(p.Name) - 1]
	fmt.Printf("My name is %s, shortName is %s", p.Name, p.shortName)
}

func (p Person) hello() {
	fmt.Println("Hello")
}