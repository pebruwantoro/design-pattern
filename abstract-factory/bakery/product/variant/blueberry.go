package variant

import (
	"abstract-factory/bakery/product"
	"abstract-factory/bakery/product/cake"
	"abstract-factory/bakery/product/croissant"
	"abstract-factory/bakery/product/donut"
)

type Blueberry struct{}

func (c *Blueberry) MakeDonut() donut.IDonut {
	return &product.BlueberryDonut{
		Donut: donut.Donut{
			Name:  "Blueberry Donut",
			Price: 15000.00,
		},
	}
}

func (c *Blueberry) MakeCake() cake.ICake {
	return &product.BlueberryCake{
		Cake: cake.Cake{
			Name:  "Blueberry Cake",
			Price: 200000.00,
		},
	}
}

func (c *Blueberry) MakeCroissant() croissant.ICroissant {
	return &product.BlueberryCroissant{
		Croissant: croissant.Croissant{
			Name:  "Blueberry Croisant",
			Price: 10000.00,
		},
	}
}

// func (s *Blueberry) MakeCookies() cookies.ICookies {
// 	return &product.BlueberryCookies{
// 		Cookies: cookies.Cookies{
// 			Name:  "Blueberry Cookies",
// 			Price: 19200.50,
// 		},
// 	}
// }
