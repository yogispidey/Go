package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	loc := sqrtError{
		lat:  "50.2289 N",
		long: "99.4656 W",
		err:  fmt.Errorf("cannot use number less than zero"),
	}
	_, err := sqrt(-10.23, loc)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64, err error) (float64, error) {
	if f < 0 {
		// write your error code here
		return 0, err
	}
	return 42, nil
}
