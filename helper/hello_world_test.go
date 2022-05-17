package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i:=0; i < b.N; i++ {
		HelloWorld("Fahmi")
	}
}

func BenchmarkHelloWorld2(b *testing.B) {
	for i:=0; i < b.N; i++ {
		HelloWorld("Fahmi2")
	}
}

func BenchmarkTableTest(b *testing.B) {

	tests := []struct {
		name string
		request string
		expected string
	} { 
		{
		name : "Fahmi",
		request : "Fahmi",
		expected : "Hello Fahmi",
	}, {
		name : "Firdaus",
		request : "Firdaus",
		expected : "Hello Firdaus",
	},
}
for _ , test := range tests {
	b.Run(test.name , func(b *testing.B) {
		for i:=0; i < b.N; i++ {
			HelloWorld(test.name)
		}
		
	})
}	

}

func TestTabletest (t *testing.T) {

	tests := []struct {
		name string
		request string
		expected string
	} { 
		{
		name : "Fahmi",
		request : "Fahmi",
		expected : "Hello Fahmi",
	}, {
		name : "Firdaus",
		request : "Firdaus",
		expected : "Hello Firdaus",
	},
}

for _ , test := range tests {
	t.Run(test.name , func(t *testing.T) {
		result := HelloWorld(test.request)
		assert.Equal(t , result , test.expected)
	})
}

}

// func TestMain (m *testing.M) {
// 	fmt.Println("BEFORE UNIT TEST")
// 	m.Run()
// 	fmt.Println("AFTER UNIT TEST")
// }

func TestSubTest (t *testing.T) {
	t.Run("Fahmi" , func(t *testing.T) {
		result := HelloWorld("Fahmi")
		assert.Equal(t , "Hello Fahmi" , result , "they should be equal")
		
	})

	t.Run("Firdaus" , func(t *testing.T) {
		result := HelloWorld("Firdaus")
		assert.Equal(t , "Hello Firdaus" , result , "they should be equal")
	})
}

func TestHelloWorldAssert (t *testing.T) {

result := HelloWorld("Fahmi")
assert.Equal(t , "Hello Fahmi" , result , "they should be equal")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fahmi")
	if result != "Hello Fahmi" {
		t.Error("Error must be Hello Fahmi")
	}
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Fahmi2")
	if result != "Hello Fahmi2" {
		t.Error("Error must be Hello Fahmi2")
	}
}