package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func DoDelete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to delete user %v\n", id)

	// start of protected code changes
	lock.Lock()
	var tmp = []*User{}
	for _, u := range Db {
		if u.Id == id {
			continue
		}
		tmp = append(tmp, u)
	}
	Db = tmp
	// end of protected code changes
	lock.Unlock()

}
