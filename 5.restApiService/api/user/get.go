package user

import (
	"encoding/json"
	"net/http"
)

func DoGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	je := json.NewEncoder(w)
	je.Encode(Db)
}
