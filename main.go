package main

import (
	"fmt"
	abstractfactory "gminato/design-pattern-in-go/abstract-factory"
	"gminato/design-pattern-in-go/factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(musket)
	printDetails(ak47)

	// Abstract factory pattern incomplete

	shoeFactor := abstractfactory.GetSportFactory("Nike")
	nikeShoe := shoeFactor.MakeShoe()
	nikeShoe.SetSize(15)
	fmt.Println(nikeShoe.GetSize())

	adidasFactory := abstractfactory.GetSportFactory("Adidas")
	AdidasShoe := adidasFactory.MakeShoe()
	fmt.Print(AdidasShoe.GetSize())
	nikeShirt := shoeFactor.MakeShirt()
	fmt.Println(nikeShirt.GetType())
	fmt.Println(adidasFactory.MakeShirt().GetType())
}
func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
