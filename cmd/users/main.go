package main

import(
	"fmt"
	"github.com/golang-architechture"
	

) 

func main() {
	storage := mongo{}
	p1 := Person{first: "Daniel"}
	p2 := Person{first: "Vaishali"}

	put(storage, 0, p1)
	put(storage, 1, p2)

	fmt.Println(get(storage, 0))
	fmt.Println(get(storage, 1))

}
