package main

import (
	"fmt"
	"github.com/kd2718/gomarshal/person"
	"encoding/json"
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

	encoded, err := json.Marshal(kory)
	if err != nil {
		fmt.Println("Problem")
	}

	e2, err := kory.Marshal()


	fmt.Println("Person encoded", string(encoded), string(e2))

	var newperson, n2 person.Person

	_ = json.Unmarshal(encoded, &newperson)
	_ = n2.Unmarshal(e2)

	fmt.Println("Person decoded", newperson, n2)

	devencoded, err := json.Marshal(nando)
	dev2, err := nando.Marshal()

	fmt.Println("dev json marshal test", string(devencoded), string(dev2))

	var newdev, newdev2 person.Developer

	_ = json.Unmarshal(devencoded, &newdev)
	_ = newdev2.Unmarshal(dev2)
	fmt.Println("newdev2 ok")

	fmt.Println("dev json unmarshal dev test", newdev, newdev2)

}
