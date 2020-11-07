//Print out the remainder (modulus) for the numbers between 10 and 100 when it is divided by 4
//Modulus or remainder operator is just %
//Divide is just /

package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
	}
}
