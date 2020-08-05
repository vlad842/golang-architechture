package main

import(
	"fmt"
	"github.com/golang-architechture"
	"github.com/golang-architechture/storage/mongo"
	
) 

func main() {
	storage := mongo.Db{}
	p1 := architecture.Person{First: "Daniel"}
	p2 := architecture.Person{First: "Vaishali"}
	
	architecture.Put(storage,0,p1)
	architecture.Put(storage, 1, p2)

	fmt.Println(architecture.Get(storage, 0))
	fmt.Println(architecture.Get(storage, 1))

}
