package main

import "fmt"

type User struct {
	Name  string
	Age   uint8
	Email string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}

func main() {
	var count int = 20
	fmt.Println(&count)

	user := &User{
		Name:  "Tejaswi",
		Age:   24,
		Email: "kasattejasvi@gmail.com",
	}
	fmt.Println(user.GetName())
	fmt.Println(user.GetEmail())
	user.SetName("Tejaswi Kasat")
	fmt.Println(*user)
}
