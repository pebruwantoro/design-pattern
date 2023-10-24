package variant

import (
	"abstract-factory/bakery/product"
	"abstract-factory/bakery/product/cake"
	"abstract-factory/bakery/product/croissant"
	"abstract-factory/bakery/product/donut"
)

type Strawberry struct{}

func (s *Strawberry) MakeDonut() donut.IDonut {
	return &product.StrawberryDonut{
		Donut: donut.Donut{
			Name:  "Strawberry Donut",
			Price: 15000.00,
		},
	}
}

func (s *Strawberry) MakeCake() cake.ICake {
	return &product.StrawberryCake{
		Cake: cake.Cake{
			Name:  "Strawberry Cake",
			Price: 200000.00,
		},
	}
}

func (s *Strawberry) MakeCroissant() croissant.ICroissant {
	return &product.StrawberryCroissant{
		Croissant: croissant.Croissant{
			Name:  "Strawberry Croissant",
			Price: 10000.00,
		},
	}
}

// func (s *Strawberry) MakeCookies() cookies.ICookies {
// 	return &product.StrawberryCookies{
// 		Cookies: cookies.Cookies{
// 			Name:  "Strawberry Cookies",
// 			Price: 19200.50,
// 		},
// 	}
// }
