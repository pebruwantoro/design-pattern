package bakery

import (
	"abtract-factory/bakery/product/cake"
	"abtract-factory/bakery/product/croissant"
	"abtract-factory/bakery/product/donut"
	"abtract-factory/bakery/product/variant"
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
