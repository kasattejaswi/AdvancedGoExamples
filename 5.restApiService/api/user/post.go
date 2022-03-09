package user

import (
	"encoding/json"
	"log"
	"net/http"
)

func DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jd := json.NewDecoder(r.Body)
	aUser := &User{}
	err := jd.Decode(aUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start of protected code changes
	lock.Lock()
	nextUserId++
	aUser.Id = nextUserId
	Db = append(Db, aUser)
	// end of protected code changes
	lock.Unlock()

	respUser := User{Id: aUser.Id, Username: aUser.Username, Password: aUser.Password}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}
