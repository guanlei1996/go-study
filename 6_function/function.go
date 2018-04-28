package main

import (
	"github.com/guanlei1996/go-study/5_function/entity"
	"fmt"
)

func main() {
	tom := entity.Person{Name: "tom", Age: 12}
	fmt.Println("call modify method before : " + tom.Name)
	tom.MethodWithValue("Lay")
	fmt.Printf("value method after %s\n", tom.Name)
	tom.MethodWithPointer("Lay")
	fmt.Println("pointer method after : " + tom.Name)

	jerry := entity.Person{Name: "Jerry"}
	fmt.Println("call modify function before : " + tom.Name)
	entity.FunctionWithValue(jerry, "Lay")
	fmt.Printf("value function after %s\n", tom.Name)
	entity.FunctionWithPointer(&jerry, "Lay")
	fmt.Println("pointer method function : " + tom.Name)
}
