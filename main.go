package main

import (
	"fmt"
	abstractfactory "gminato/design-pattern-in-go/abstract-factory"
	"gminato/design-pattern-in-go/builder"
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

	// builder pattern

	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}
func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
