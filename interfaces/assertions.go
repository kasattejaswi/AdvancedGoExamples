package main

import (
	"fmt"
	"reflect"
)

type Registry struct {
	register interface{}
	alias    string
}

func Exec(registry interface{}, args ...interface{}) {
	s, ok := registry.(Registry)
	if !ok {
		fmt.Println("Error while asserting Registry")
		return
	}
	fmt.Println("First assertion completed")
	f, ok := s.register.(func() int)
	if !ok {
		fmt.Println("Error while asserting function")
		return
	}
	fmt.Println("Second assertion completed")
	result := f()
	fmt.Println(result)
}

func Add() int {
	return 10 + 20
}

func sub(a, b int) int {
	return a - b
}

func main() {
	registry := Registry{
		register: Add,
		alias:    "Add",
	}
	Exec(registry, 10, 20)
	fmt.Println(reflect.TypeOf(sub))
}
