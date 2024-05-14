package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) getHouse() House {
	return House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}
}

// setDoorType implements IHouseBuilder.
func (i *IglooBuilder) setDoorType() {
	i.doorType = "no doors"
}

// setFloor implements IHouseBuilder.
func (i *IglooBuilder) setFloor() {
	i.floor = 0
}

// setWindowType implements IHouseBuilder.
func (i *IglooBuilder) setWindowType() {
	i.windowType = "ice window"
}
