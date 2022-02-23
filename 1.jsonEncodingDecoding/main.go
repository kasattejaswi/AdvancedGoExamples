package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserDb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func main() {
	// Create json with data
	users := []User{
		{
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
		{
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
		{
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
	}
	db := UserDb{
		Users: users,
		Type:  "Simple",
	}
	createJsonFile(db)

	// Read json data
	dbR := &UserDb{}
	readJsonFile(dbR, "users.db.json")
	fmt.Println(dbR)

}

func createJsonFile(db UserDb) {
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)
	f, err := os.Create("users.db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}

func readJsonFile(db *UserDb, fileName string) {
	f, err := os.Open("users.db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)

	dec.Decode(&db)
}
