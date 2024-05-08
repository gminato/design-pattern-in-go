package abstractfactory

type IShoe interface {
	SetSize(size int)
	GetSize() int
}

type Shoe struct {
	size int
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}
