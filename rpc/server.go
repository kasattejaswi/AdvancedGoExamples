package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

var database []Item

type API int

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	for _, item := range database {
		if item.Title == title {
			*reply = item
		}
	}
	return nil
}

func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	for i, v := range database {
		if v.Title == edit.Title {
			database[i] = edit
			*reply = edit
		}
	}
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	for i, v := range database {
		if v.Title == item.Title {
			database = append(database[:i], database[i+1:]...)
			*reply = v
		}
	}
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API: ", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("error listening: ", err)
	}
	log.Printf("Serving RPC on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}
	// fmt.Println("Initial database", database)
	// a := Item{title: "first", body: "a first item"}
	// b := Item{title: "second", body: "a second item"}
	// c := Item{title: "third", body: "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println("Second database", database)

	// DeleteItem(b)
	// fmt.Println("Third database", database)

	// EditItem("third", Item{title: "fourth", body: "a new item"})
	// fmt.Println("Fourth database", database)

	// x := GetByName("first")
	// y := GetByName("fourth")
	// fmt.Println(x, y)
}
