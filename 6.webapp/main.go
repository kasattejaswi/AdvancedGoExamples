package main

import (
	"log"
	"net/http"

	"github.com/kasattejaswi/restapiservice/api/user"
)

func main() {
	//register static files handler 'index.html' -> /client/index.html
	http.Handle("/", http.FileServer(http.Dir("./client")))

	// register rest api handler
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":5001", nil))
}
