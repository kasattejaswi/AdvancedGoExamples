package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var counter int

type (
	Currency float64
	Item     struct {
		Name     string   `json:"name,omitempty"`
		Quantity int      `json:"quantity,omitempty" json:"quantity,omitempty"`
		Price    Currency `json:"price,omitempty" json:"price,omitempty"`
	}
	Store struct {
		Items map[string]Item `json:"items,omitempty"`
	}
	Dashboard struct {
		Users         uint   `json:"users,omitempty"`
		UsersLoggedIn uint   `json:"users_logged_in,omitempty"`
		Inventory     *Store `json:"inventory"`
	}
)

var dasboard chan *Dashboard

func main() {
	dasboard = make(chan *Dashboard)
	go updateDashboard()
	//register static files handler 'index.html' -> /client/index.html
	http.Handle("/", http.FileServer(http.Dir("./client")))

	// register rest api for server sent event /sse/dashboard
	http.HandleFunc("/sse/dashboard", dashboardHandler)
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(<-dasboard)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
	fmt.Printf("data: %v\n\n", buf.String())
}

func updateDashboard() {
	for {
		inv := updateInventory()
		db := &Dashboard{
			Users:         uint(rand.Uint32()),
			UsersLoggedIn: uint(rand.Uint32() % 200),
			Inventory:     inv,
		}
		dasboard <- db
	}
}

func updateInventory() *Store {
	for {
		inv := &Store{}
		inv.Items = make(map[string]Item)
		a := Item{Name: "Books", Price: 33.59, Quantity: int(rand.Int31() % 53)}
		inv.Items["book"] = a
		a = Item{Name: "Bicycles", Price: 190.89, Quantity: int(rand.Int31() % 232)}
		inv.Items["bicycle"] = a
		a = Item{Name: "Water Bottles", Price: 10.02, Quantity: int(rand.Int31() % 93)}
		inv.Items["wbottle"] = a
		a = Item{Name: "RC Car", Price: 83.19, Quantity: int(rand.Int31() % 73)}
		inv.Items["rccar"] = a
		return inv
	}
}
