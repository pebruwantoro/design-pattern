package variant

import (
	"abstract-factory/bakery/product"
	"abstract-factory/bakery/product/cake"
	"abstract-factory/bakery/product/croissant"
	"abstract-factory/bakery/product/donut"
)

// Concrete Factory
type Chocolate struct{}

func (c *Chocolate) MakeDonut() donut.IDonut {
	return &product.ChocolateDonut{
		Donut: donut.Donut{
			Name:  "Chocolate Donut",
			Price: 15000.00,
		},
	}
}

func (c *Chocolate) MakeCake() cake.ICake {
	return &product.ChocolateCake{
		Cake: cake.Cake{
			Name:  "Chocolate Cake",
			Price: 200000.00,
		},
	}
}

func (c *Chocolate) MakeCroissant() croissant.ICroissant {
	return &product.ChocolateCroissant{
		Croissant: croissant.Croissant{
			Name:  "Chocolate Croisant",
			Price: 10000.00,
		},
	}
}

// func (s *Chocolate) MakeCookies() cookies.ICookies {
// 	return &product.ChocolateCookies{
// 		Cookies: cookies.Cookies{
// 			Name:  "Chocolate Cookies",
// 			Price: 19200.50,
// 		},
// 	}
// }
