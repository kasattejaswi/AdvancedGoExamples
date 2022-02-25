package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

// Client example using context, see server code for server example

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	// defer cancel()
	req, err := http.NewRequest("GET", "http://localhost:9898", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal("Status code was not OK:", res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}
