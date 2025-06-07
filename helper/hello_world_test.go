package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Mustofa Adny")

	if result != "Hello Mustofa Adny" {
		t.Fail() // tetap menjalankan kode di bawahnya
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Ucok Udin")

	if result != "Hello Ucok Udin" {
		t.FailNow() // tidak menjalankan kode di bawahnya
	}

	fmt.Println("TestHelloWorld2 done")
}