package product

import (
	"abtract-factory/bakery/product/cake"
	"abtract-factory/bakery/product/croissant"
	"abtract-factory/bakery/product/donut"
)

// Concrete Product
type ChocolateDonut struct {
	donut.Donut
}

type ChocolateCake struct {
	cake.Cake
}

type ChocolateCroissant struct {
	croissant.Croissant
}

type StrawberryDonut struct {
	donut.Donut
}

type StrawberryCake struct {
	cake.Cake
}

type StrawberryCroissant struct {
	croissant.Croissant
}
