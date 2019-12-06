package main

import (
	"fmt"
	"oop/oopMid/oopEncapsulation/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p.GetAge())
}

