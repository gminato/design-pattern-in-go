package abstractfactory

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShirt() Ishirt
}

func GetSportFactory(factoryName string) ISportFactory {
	if factoryName == "Nike" {
		return &Nike{}
	}
	if factoryName == "Adidas" {
		return &Adidas{}
	}
	return nil
}
