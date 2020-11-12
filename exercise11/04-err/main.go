package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v ", se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		ctn := fmt.Errorf("Can't take sqrt of negative number: %v ", f)
		return 0, sqrtError{ctn}

	}
	return 42, nil
}

// Starting with this code use the sqrt.Error struct as a value of type error
