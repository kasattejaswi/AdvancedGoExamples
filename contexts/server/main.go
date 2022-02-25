package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Server code using context and cancellation

type key int

const requestIdKey = key(42)

func main() {
	http.HandleFunc("/", Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:9898", nil))
}

// Custom log print
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		log.Println("could not find request id in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

// modify http.HandlerFunc to include ID on every call
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIdKey, id)
		f(w, r.WithContext(ctx))
	}
}

// Request handler implementation
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// ctx = context.WithValue(ctx, int(42), int64(100))
	Println(ctx, "Handler started")
	defer Println(ctx, "Handler ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprint(w, "Hello")
	case <-ctx.Done():
		Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
