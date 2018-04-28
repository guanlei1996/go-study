package entity

import "fmt"

type User struct {
	Name string
}

func (u User)SayHello() {
	fmt.Printf("Hello, I'm %s, My role is user\n", u.Name)
}
