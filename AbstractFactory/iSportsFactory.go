package AbstractFactory

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSportsFactory(brand string) ISportsFactory {
	switch brand {
	case "adidas":
		return &AdidasFactory{}

	case "nike":
		return &NikeFactory{}

	}
	return nil
}
