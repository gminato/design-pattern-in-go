package abstractfactory

type Nike struct {
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			size: 12,
		},
	}
}
