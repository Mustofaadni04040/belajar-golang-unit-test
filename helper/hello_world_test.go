package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Mustofa Adny")

	if result != "Hello Mustofa Adny" {
		panic("Result is not Hello Mustofa Adny")
	}
}