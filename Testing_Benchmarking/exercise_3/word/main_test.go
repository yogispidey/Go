package word

import (
	"Go/Testing_Benchmarking/exercise_3/quote"
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("expected:", 3, "got:", n)
	}
}
