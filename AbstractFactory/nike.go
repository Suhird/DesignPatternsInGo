package main

// Concrete Nike Product factory

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 15,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 15,
		},
	}
}
