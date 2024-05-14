package builder

type Director struct {
	builder IHouseBuilder
}

func NewDirector(b IHouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setFloor()
	return d.builder.getHouse()
}
