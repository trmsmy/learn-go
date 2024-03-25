package main

import "fmt"

type gasEngine struct {
	name    string
	mpg     uint16
	gallons uint16
	manufacturer
	int
}

type electricEngine struct {
	name  string
	kwh   uint8
	mpwkh uint8
	manufacturer
}

func (g gasEngine) milesLeft() uint16 {
	return g.gallons * g.mpg
}

func (e electricEngine) milesLeft() uint16 {
	return uint16(e.kwh) * uint16(e.mpwkh)
}

type manufacturer struct {
	manufacturer string
}

type engine interface {
	milesLeft() uint16
}

func canMakeIt(engine engine, miles uint16) {
	if miles <= engine.milesLeft() {
		fmt.Println("You can make it there !")
	} else {
		fmt.Printf("your %T Need more fuel to get there ! \n", engine)
	}
}

func main() {
	fmt.Println("types")
	var gasEng gasEngine = gasEngine{"V6", 20, 15, manufacturer{"Engine LLC"}, 10}

	var anonymous = struct {
		f1 int
		f2 string
	}{10, "user1"}

	fmt.Printf("ID=%v and Name=%v \n", anonymous.f1, anonymous.f2)

	fmt.Printf("myengine: %v, milesLeft=%v \n", gasEng, gasEng.milesLeft())

	canMakeIt(gasEng, 200)

	var elecEngine electricEngine = electricEngine{"Vroom", 25, 15, manufacturer{"Engine LLC"}}

	fmt.Println(elecEngine)

	canMakeIt(elecEngine, 400)
}
