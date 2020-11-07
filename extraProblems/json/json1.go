package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Geralt",
		Last:  "Of Rivia",
		Age:   90,
	}
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
