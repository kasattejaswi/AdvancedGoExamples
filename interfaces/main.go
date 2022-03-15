package main

import (
	"fmt"
)

type Greeting interface {
	Greet()
}

type name struct {
	first string
	last  string
}

//func assert(g Greeting) {
//	switch g.(type) {
//	case name:
//		v, ok := g.(name)
//		if !ok {
//			fmt.Println("Failed to assert type name")
//			return
//		}
//		fmt.Println("First name:", v.first, "Last name:", v.last)
//	case mood:
//		v, ok := g.(mood)
//		if !ok {
//			fmt.Println("Failed to assert type mood")
//			return
//		}
//		fmt.Println("Mood is", v.name)
//	}
//}

func assert(g Greeting) {
	switch g.(type) {
	case name:
		s, ok := g.(name)
		if !ok {
			fmt.Println("name type is not ok")
			return
		}
		fmt.Printf("First name is %v and last name is %v\n", s.first, s.last)
	case mood:
		s, ok := g.(mood)
		if !ok {
			fmt.Println("mood type is not ok")
			return
		}
		fmt.Printf("The mood is %v\n", s.name)
	default:
		fmt.Println("The type is unknown")
	}
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
	assert(n)
	assert(m)
}
