package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/starting-code/dog"
	"testing"
)

func BenchmarkYears(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(10)
	}
}

func ExampleYears()  {
	fmt.Println(dog.Years(10))
	//Output:
	//70
}

func ExampleYearsTwo(){
	fmt.Println(dog.YearsTwo(10))
	//Output:
	//70
}

func TestYears(t *testing.T)  {
	n := dog.Years(10)
	if n != 70 {
		t.Error("Expected", 70, "Got", n)
	}
}

func TestYearsTwo(t *testing.T)  {
	n := dog.YearsTwo(10)
	if n != 70 {
		t.Error("Expected", 70, "Got", n)
	}
}