package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("before unit test")
	m.Run()
	// after
	fmt.Println("after unit test")
}

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

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Alexander")

	assert.Equal(t, "Hello Alexander", result, "Result must be Hello Alexander") // Fail() if error

	fmt.Println("TestHelloWorldAssertion done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bambang")
	
	require.Equal(t, "Hello Bambang", result, "Result must be Hello Bambang") // FailNow() if error

	fmt.Println("TestHelloWorldRequire done")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on mac os")
	}

	result := HelloWorld("Adipati")
	require.Equal(t, "Hello Adipati", result, "Result must be Hello Adipati")
}

func TestSubTest(t *testing.T) {
	t.Run("Subtest 1", func(t *testing.T) {
		result := HelloWorld("Memet")
		require.Equal(t, "Hello Memet", result, "Result must be Hello Memet")
	})

	t.Run("Subtest 2", func(t *testing.T) {
		result := HelloWorld("BiBi")
		require.Equal(t, "Hello BiBi", result, "Result must be Hello BiBi")
	})
}

type Hello struct {
	name 		string
	request 	string
	expected 	string
}

func TestHelloWorldWithTable(t *testing.T) {
	tests := []Hello {
		{
			name: "Test table 1",
			request: "Siti",
			expected: "Hello Siti",
		},
		{
			name: "Test table 2",
			request: "Subandi",
			expected: "Hello Subandi",
		},
		{
			name: "Test table 3",
			request: "Sujarwo",
			expected: "Hello Sujarwo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be " + test.expected)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ { // atau shortcut: b.loop
		HelloWorld("Mustofa Adny")
	}
}