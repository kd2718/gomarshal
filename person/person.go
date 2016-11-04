/*
Person and Developer structs. To be used in further packages
*/
package person

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Age int

func (a *Age) Birthday() {
	*a++ // Question: Is this the "proper" way to do an operations like this?
}

type Person struct {
	Age      Age
	First    string
	Last     string
	CoolInfo interface{}
	personChan chan Person
}

func (p *Person) Loop(outside chan Person) {
	killout := false
	guy := new(Person)
	// register everyone
	select {
	case *guy = <-outside:
		fmt.Println(p.First, "outside", guy.First)
		guy.personChan <- *p
	case *guy = <-p.personChan:
		fmt.Println(p.First, "self", guy.First)
	}

    fmt.Println("I am", p.First, p.Age, "he is", guy.First, guy.Age)
	AgeLoop:
	for ; p.Age < Age(200); {
		fmt.Println("I am", p.First, p.Age, "he is", guy.First, guy.Age)
		select {
		case _, ok := <-outside:
		if ok {
			killout = false
		} else {
			killout = true
		}
			break AgeLoop
		case *guy = <- p.personChan:
			if p.Age < guy.Age {
				fmt.Println("I am", p.First, p.Age, "he is", guy.First, guy.Age)
				fmt.Println("falling behind, extra birthdays for", p.First)
				p.Birthday()
				p.Birthday()
				p.Birthday()
			}
			guy.personChan <- *p
		default:
			guy.personChan <- *p
		}
		p.Birthday()
	}
	_ = killout
	//if killout {
	//	close(outside)
	//} else {
	//	outside <- *p
	//}
	fmt.Println(p.First, "Is done...", p.Age)
	return
}

type Looper interface{
	Loop(c chan Person)
}

func (p Person) String() string {
	return fmt.Sprintf("Full name: %v, %v  Age: %v, Cool Info: %v", p.First, p.Last, p.Age, p.CoolInfo)
}

func (p *Person) Birthday() {
	person := *p
	person.Age.Birthday()
	*p = person
}

func (p Person) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (d Developer) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func typeSwitch(v interface{}) interface{} {
	switch arg := v.(type) {
	default:
		fmt.Println("Couldn't match type", arg)
		return v
	//case int, float32, float64:
	//	fmt.Println("standard")
	//	return Age(v)
	case int:
		return v.(int)
	case float32:
		return int(v.(float32))

	case float64:
		return int(v.(float64))
	case string:
		return v
	case map[string]interface{}:
		fmt.Println("Map")
		return v.(map[string]interface{})
	case []interface{}:
		fmt.Println("Slice", reflect.TypeOf(arg))
		return arg

	}
}

func (p *Person) Unmarshal(b []byte) (err error) {
	//var f interface{}
	err = json.Unmarshal(b, &p)
	// type switch here


	//m := f.(map[string]interface{})
	//for key, value := range m {
	//	m[key] = typeSwitch(value)
	//}
	//
	//p.Age = Age(m["Age"].(int))
	//p.First = m["First"].(string)
	//p.Last = m["Last"].(string)
	//p.CoolInfo = m["CoolInfo"]

	return
}

func (d *Developer) Unmarshal(b []byte) (err error) {
	var per Person
	err = per.Unmarshal(b)

	var f, ok interface{}
	_ = ok
	err = json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	for key, value := range m {
		m[key] = typeSwitch(value)
	}

	d.Person = &per

	d.Language = m["Language"].(string)

	return
}

type Developer struct {
	*Person
	Language string
}

func (d Developer) String() string {
	return fmt.Sprintf("%v, uses language: %v", d.Person, d.Language)
}

// Developers age twice as fast
func (d *Developer) Birthday() {
	dev := *d
	dev.Person.Birthday()
	dev.Person.Birthday()
	*d = dev
}

type Ager interface {
	Birthday()
}

func HardTimes(a Ager) {
	fmt.Println("Hard times have fallen...")
	a.Birthday()
	a.Birthday()
}

func NewPerson(age Age, first, last string, v interface{}) *Person {
	return &Person{
		age,
		first,
		last,
		v,
		make(chan Person),
	}
}

func NewDeveloper(age Age, first, last, language string, v interface{}) *Developer {
	return &Developer{
		NewPerson(
			age,
			first,
			last,
			v,
		),
		language,
	}
}
