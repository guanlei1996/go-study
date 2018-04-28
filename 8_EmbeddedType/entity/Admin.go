package entity

import "fmt"

type Admin struct {
	User
	Level string
}

func (a Admin)SayHello()  {
	fmt.Printf("Hello, I'm %s, My role is admin, My level is %s\n", a.Name, a.Level)
}