package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func DoPut(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to update user %v\n", id)

	jd := json.NewDecoder(r.Body)
	aUser := &User{}
	err = jd.Decode(aUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// start of protected code changes
	lock.Lock()
	for i, u := range Db {
		if u.Id == id {
			if aUser.Username != "" {
				Db[i].Username = aUser.Username
			}
			if aUser.Password != "" {
				Db[i].Password = aUser.Password
			}
		}
	}
	// end of protected code changes
	lock.Unlock()
}
