package main

import "fmt"

type Greeting interface {
	Greet()
}

type name struct {
	first string
	last  string
}

func (n name) Greet() {
	fmt.Println("Hello", n.first, n.last, ", Welcome to Interfaces")
}

type mood struct {
	name string
}

func (m mood) Greet() {
	fmt.Println("Your mood is", m.name)
}

func main() {
	n := name{"John", "Doe"}
	m := mood{"Happy"}
	n.Greet()
	m.Greet()
}
