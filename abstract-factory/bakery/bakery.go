package bakery

import (
	"abstract-factory/bakery/product/cake"
	"abstract-factory/bakery/product/croissant"
	"abstract-factory/bakery/product/donut"
	"abstract-factory/bakery/product/variant"
	"errors"
)

// Abstract Factory
type IMamatosaiBakery interface {
	MakeDonut() donut.IDonut
	MakeCake() cake.ICake
	MakeCroissant() croissant.ICroissant
	// MakeCookies() cookies.ICookies
}

func GetBakeryMenu(menu string) (IMamatosaiBakery, error) {
	if menu == "chocolate" {
		return &variant.Chocolate{}, nil
	}
	if menu == "strawberry" {
		return &variant.Strawberry{}, nil
	}
	if menu == "blueberry" {
		return &variant.Blueberry{}, nil
	}
	// ADD THE VARIANT OF PRODUCT

	return nil, errors.New("menu not exist")
}
