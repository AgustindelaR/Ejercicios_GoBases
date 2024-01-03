package main

type Product interface {
	Price() float32
}

type Small struct {
	PriceUnit float32
}

func (s Small) Price() float32 {
	return s.PriceUnit
}

type Medium struct {
	PriceUnit float32
}

func (m Medium) Price() float32 {
	return m.PriceUnit * 1.03
}

type Large struct {
	PriceUnit float32
}

func (l Large) Price() float32 {
	return l.PriceUnit*1.06 + 2500
}

// patrón factory
func CreateProduct(productType string, price float32) Product {
	switch productType {
	case "Small":
		return Small{PriceUnit: price}
	case "Medium":
		return Medium{PriceUnit: price}
	case "Large":
		return Large{PriceUnit: price}
	default:
		println("Tipo de producto invalido, se tomara como small")
		return Small{}
	}
}
