// Create a slice of a slice of string. Store the data in the multi dimensional slice
// Range over the records then range over the data in each record

package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken", "Not Stirred"}
	fmt.Println(jb)
	mp := []string{"Moneypenny", "Hello James"}
	fmt.Println(mp)

	jbmp := [][]string{jb, mp}

	fmt.Println(jbmp)

	for i, jbmp1 := range jbmp {
		fmt.Println("Record: ", i)
		for j, jbmp2 := range jbmp1 {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, jbmp2)
		}
	}
}
