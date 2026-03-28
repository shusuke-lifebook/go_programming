package main

type Human interface {
	Say()
}

type Person struct {
	Name string
}

// func (p Person) Say() {
// 	fmt.Println(p.Name)
// }

func main() {
	var mike Human = Person{"Mike"} // cannot use Person{…} (value of struct type Person) as Human value in variable declaration: Person does not implement Human (missing method Say)
	mike.Say()
}
