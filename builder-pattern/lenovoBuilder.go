package builderpattern

type lenovoBuilder struct {
	ram         string
	storage     string
	processor   string
	motherboard string
}

// GetComputer implements Icomputer.
func (l lenovoBuilder) GetComputer() *Computer {
	return &Computer{
		Ram:         l.ram,
		Processor:   l.processor,
		Storage:     l.storage,
		Motherboard: l.motherboard,
	}
}

func (l *lenovoBuilder) SetMotherboard(motherboard string) {
	l.motherboard = motherboard
}

func (l *lenovoBuilder) SetProcessor(processor string) {
	l.processor = processor
}

func (l *lenovoBuilder) SetRam(ramsize string) {
	l.ram = ramsize
}

func (l *lenovoBuilder) SetStorage(storage string) {
	l.storage = storage
}

func getLenovoBuilder() Icomputer {
	return &lenovoBuilder{}
}
