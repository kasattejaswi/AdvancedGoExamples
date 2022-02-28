package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type User struct {
	Id       int    `xml:"id,attr"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Email    string `xml:"email"`
}

type UserDb struct {
	XMLName xml.Name `xml:"userdb"`
	Type    string   `xml:"type"`
	Users   []User   `xml:"users>user"`
}

const xmlFile = "userdb.xml"

func main() {
	users := []User{
		{
			Id:       1,
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
		{
			Id:       2,
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
		{
			Id:       3,
			Username: "Tejaswi Kasat",
			Password: "changeit",
			Email:    "kasatejaswi@gmail.com",
		},
	}
	db := UserDb{
		Users: users,
		Type:  "Simple",
	}
	createXmlFile(db, xmlFile)
	var rdb UserDb
	readXmlFile(&rdb, xmlFile)
	fmt.Println(rdb)
}

func createXmlFile(db UserDb, xmlFile string) {
	f, err := os.Create(xmlFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	xmlEnc := xml.NewEncoder(f)
	xmlEnc.Encode(db)
}

func readXmlFile(db *UserDb, fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dec := xml.NewDecoder(f)

	dec.Decode(&db)
}
