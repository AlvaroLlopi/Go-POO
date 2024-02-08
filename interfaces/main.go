package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreetByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s", p.Name)
}

func (p Person) String() string {
	return "Hola soy una Persona y mi nombre es: " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s", t)
}
func (t Text) Bye() {
	fmt.Printf("Adios soy %s", t)
}

func All(gbs ...GreetByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	p := Person{Name: "Alvaro"}
	//(var t Text = "Daisy"
	//All(p, t)
	fmt.Println(p)
}
