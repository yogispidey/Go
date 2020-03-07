package main

import "testing"

func main() {

	TestSum()
}

func TestSum(t *testing.T){
	v := Sum(1,2)
	if v != 3 {
		t.Error("Expected 3 got:",v)
	}
}

func Sum(a int, b int) int  {
	return a - b
}
