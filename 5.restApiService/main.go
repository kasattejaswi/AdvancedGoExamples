package main

import (
	"log"
	"net/http"

	"github.com/kasattejaswi/restapiservice/api/user"
)

func main() {
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":5001", nil))
}
