package user

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UserAPI struct{}
type User struct {
	Id       uint64
	Username string
	Password string
}

var Db = []*User{}
var nextUserId uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
	switch r.Method {
	case http.MethodGet:
		DoGet(w, r)
	case http.MethodPost:
		DoPost(w, r)
	case http.MethodDelete:
		DoDelete(w, r)
	case http.MethodPut:
		DoPut(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method: %s on %v\n", r.Method, r.URL)
		log.Printf("Unsupported method: %v to %v\n", r.Method, r.URL)
	}
}
