package builderpattern

type Icomputer interface {
	SetRam(ramsize string)
	SetStorage(storage string)
	SetProcessor(processor string)
	SetMotherboard(motherboard string)
	GetComputer() *Computer
}

func GetBuilder(companyName string) Icomputer {
	if companyName == "lenovo" {
		return getLenovoBuilder()
	} else {
		return nil
	}
}
