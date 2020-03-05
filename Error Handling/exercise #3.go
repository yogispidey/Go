package main

import "fmt"

type customErr struct {
	e string
}

func (ce customErr) Error() string  {
	//fmt.Println(ce.e)
	return fmt.Sprintf("here is the err: %v",ce.e)
}

func foo(e error)  {
	fmt.Println(e)
	//fmt.Println(e.(customErr).e)
}

func main() {
	er := customErr{e: "This will not work"}
	foo(er)
}
