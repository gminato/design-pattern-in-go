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

func (n *Nike) MakeShirt() Ishirt {
	return &NikeShirt{
		Shirt: Shirt{
			shirtType: "Tshirt",
		},
	}
}
