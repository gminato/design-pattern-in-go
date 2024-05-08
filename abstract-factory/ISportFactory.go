package abstractfactory

type ISportFactory interface {
	MakeShoe() IShoe
}

func GetSportFactory(factoryName string) ISportFactory {
	if factoryName == "Nike" {
		return &Nike{}
	}
	return nil
}
