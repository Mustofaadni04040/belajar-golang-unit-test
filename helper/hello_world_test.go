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
func TestHelloWorld3(t *testing.T) {
	result := HelloWorld("Siti Anjani")

	if result != "Hello Siti Anjani" {
		t.Error("Result must be Hello Siti Anjani") // tetap menjalankan kode di bawahnya
	}

	fmt.Println("TestHelloWorld3 done")
}

func TestHelloWorld4(t *testing.T) {
	result := HelloWorld("Asep Anjani")

	if result != "Hello Asep Anjani" {
		t.Fatal("Result must be Hello Asep Anjani") // tidak menjalankan kode di bawahnya
	}

	fmt.Println("TestHelloWorld4 done")
}