// using the following operators, write expressions and assign their values to variables
// g ==, h <=, i >=, j !=, k <, l >

package main

import (
	"fmt"
)

func main() {
	g := (48 == 20)
	h := (48 <= 20)
	i := (48 >= 20)
	j := (48 != 20)
	k := (48 < 20)
	l := (48 > 20)
	fmt.Println(g, h, i, j, k, l)
}
