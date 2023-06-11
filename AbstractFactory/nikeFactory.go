package AbstractFactory

type NikeFactory struct {
}

func (nf *NikeFactory) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			name:  "shoe",
			size:  12,
			brand: "nike",
		},
	}
}

func (nf *NikeFactory) MakeShort() IShort {
	return &nikeShort{
		short: short{
			name:  "short",
			size:  11,
			brand: "nike",
		},
	}
}
