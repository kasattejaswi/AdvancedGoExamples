package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Setting up stuff for the tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Cleaning stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("This function uses stuff setup in main", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("This function also uses stuff setup in main", testTime)
}

func Test_add(t *testing.T) {
	expected := 5
	get := add(2, 3)
	if get != expected {
		t.Errorf("Test add() failed. Expected %d, got %d", expected, get)
	}
}
