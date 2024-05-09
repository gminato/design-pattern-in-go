package abstractfactory

type Ishirt interface {
	GetType() string
	SetType(shirtType string)
}

type Shirt struct {
	shirtType string
}

func (s *Shirt) GetType() string {
	return s.shirtType
}

func (s *Shirt) SetType(shirtType string) {
	s.shirtType = shirtType
}
