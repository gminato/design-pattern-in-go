package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// getHouse implements IHouseBuilder.
func (n *NormalBuilder) getHouse() House {
	return House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}

// setDoorType implements IHouseBuilder.
func (n *NormalBuilder) setDoorType() {
	n.doorType = "wooden door"
}

// setFloor implements IHouseBuilder.
func (n *NormalBuilder) setFloor() {
	n.floor = 2
}

// setWindowType implements IHouseBuilder.
func (n *NormalBuilder) setWindowType() {
	n.windowType = "glass window"
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
