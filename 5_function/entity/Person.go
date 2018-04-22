package entity

type Person struct {
	Age uint8
	Name string
}

func (person Person) MethodWithValue(name string) {
	person.Name = name
}

func (person *Person) MethodWithPointer(name string) {
	person.Name = name
}

func FunctionWithValue(person Person, name string) {
	person.Name = name
}

func FunctionWithPointer(person *Person, name string)  {
	person.Name = name
}