package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}
	tests := []test{
		{
			data:   []int{1, 4, 6, 8, 100},
			answer: 6,
		},
		{
			data:   []int{1, 4, 6, 8, 100},
			answer: 6,
		},
	}
	for _,v := range tests{
		n := CenteredAvg(v.data)
		if n != v.answer {
			t.Error("expected",v.answer, "got", n)
		}
	}
}
