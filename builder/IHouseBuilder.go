package builder

type IHouseBuilder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() House
}

func GetBuilder(builderType string) IHouseBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	} else {
		return newIglooBuilder()
	}
}
