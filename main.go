package main

import (
	"fmt"
	"github.com/kd2718/gomarshal/person"
)

func main() {
	kory := &person.Person{
		32,
		"Kory",
		"Donati",
		person.Age(5),
	}

	//var nandolist [2]int
	nandolist := [2]int{12, 33}

	nando := person.NewDeveloper(33, "Fernando", "Sanchez", "golang", nandolist)

	fmt.Println("initial check", kory, nando)

	e2, err := kory.Marshal()
	if err != nil{
		fmt.Println("error")
	}

	fmt.Println("Person encoded", string(e2))

	var newperson, newdev2 person.Person

	_ = newperson.Unmarshal(e2)

	fmt.Println("Person decoded", newperson)

	dev2, err := nando.Marshal()

	fmt.Println("dev json marshal test", string(dev2))

	_ = newdev2.Unmarshal(dev2)
	fmt.Println("newdev2 ok")

	fmt.Println("dev json unmarshal dev test", newdev2)

}
