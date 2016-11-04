package person

import "fmt"

func RunMarshalTest() {
		kory := NewPerson(
		32,
		"Kory",
		"Donati",
		Age(5),
	)

	//var nandolist [2]int
	nandolist := [2]int{12, 33}

	nando := NewDeveloper(33, "Fernando", "Sanchez", "golang", nandolist)

	fmt.Println("initial check", kory, nando)

	e2, err := kory.Marshal()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("Person encoded", string(e2))

	var newperson, newdev2 Person

	_ = newperson.Unmarshal(e2)

	fmt.Println("Person decoded", newperson)

	dev2, err := nando.Marshal()

	fmt.Println("dev json marshal test", string(dev2))

	_ = newdev2.Unmarshal(dev2)
	fmt.Println("newdev2 ok")

	fmt.Println("dev json unmarshal dev test", newdev2)
}
