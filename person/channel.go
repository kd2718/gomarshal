package person

import (
	"time"
	"fmt"
)

func RunChannelTest() {
	kory := NewPerson(33, "kory", "donati", true)
	nando := NewPerson(35, "Fernando", "Sanchez", 42)

	var k, n Looper

	k = kory
	n = nando

	out := make(chan Person)
	defer close(out)

	go k.Loop(out)
	out <- *nando
	go n.Loop(out)

	<- time.After(10 * time.Second)

	fmt.Println("Main loop exit")
}