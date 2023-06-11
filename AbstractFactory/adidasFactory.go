package AbstractFactory

type AdidasFactory struct {
}

func (af *AdidasFactory) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			name:  "shoe",
			size:  12,
			brand: "adidas",
		},
	}
}

func (af *AdidasFactory) MakeShort() IShort {
	return &adidasShort{
		short: short{
			name:  "short",
			size:  11,
			brand: "adidas",
		},
	}
}
