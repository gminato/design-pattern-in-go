package abstractfactory

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			size: 12,
		},
	}
}

func (a *Adidas) MakeShirt() Ishirt {
	return &AdidasShirt{
		Shirt: Shirt{
			shirtType: "Sports wear",
		},
	}
}
