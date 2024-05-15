package builderpattern

type Director struct {
	builder Icomputer
}

func NewDirector(b Icomputer) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b Icomputer) {
	d.builder = b
}

func (d *Director) BuildRam(ram string) *Director {
	d.builder.SetRam(ram)
	return d
}

func (d *Director) BuildStorage(storage string) *Director {
	d.builder.SetStorage(storage)
	return d
}

func (d *Director) BuildProcessor(processor string) *Director {
	d.builder.SetProcessor(processor)
	return d
}

func (d *Director) BuildMotherboard(motherboard string) *Director {
	d.builder.SetMotherboard(motherboard)
	return d
}

func (d *Director) GetComputer() *Computer {
	return d.builder.GetComputer()
}
