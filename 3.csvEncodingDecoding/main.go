package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Age      int    `json:"age"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserDb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func main() {
	// Encoding
	var users UserDb
	readJsonFile(&users, "users.db.json")

	f, err := os.Create("usersDb.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(GetHeader())

	for _, user := range users.Users {
		csv_record := user.EncodeAsStrings()
		w.Write(csv_record)
	}

	w.Flush()
	if nil != w.Error() {
		log.Fatalln(w.Error())
	}

	// Decoding
	decodeCsv("usersDb.csv")
}

func GetHeader() []string {
	return []string{
		`id`,
		`fname`,
		`lname`,
		`age`,
		`username`,
		`email`,
		`password`,
	}
}

func (u *User) EncodeAsStrings() (ss []string) {
	ss = make([]string, 0, 6)
	ss = append(ss, strconv.Itoa(u.Id))
	ss = append(ss, u.Fname)
	ss = append(ss, u.Lname)
	ss = append(ss, strconv.Itoa(u.Age))
	ss = append(ss, u.Username)
	ss = append(ss, u.Email)
	ss = append(ss, u.Password)
	return
}

func (user *User) fromCsv(ss []string) {
	if nil == user {
		return
	}
	if nil == ss {
		return
	}

	user.Id, _ = strconv.Atoi(ss[0])
	user.Fname = ss[1]
	user.Lname = ss[2]
	user.Age, _ = strconv.Atoi(ss[3])
	user.Username = ss[4]
	user.Email = ss[5]
	user.Password = ss[6]
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

func decodeCsv(filepath string) {
	log.Println("Decoding csv file")
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read() // Reading header and ignoring it
	for {
		csvRecord, err := r.Read()
		if err == nil {
			process(csvRecord)
		} else if err == io.EOF {
			break
		} else {
			log.Fatal(err)
		}
	}
}

func process(ss []string) {
	u := &User{}
	u.fromCsv(ss)
	fmt.Println(u.Id, u.Fname, u.Lname, u.Age, u.Username, u.Email, u.Password)
}
