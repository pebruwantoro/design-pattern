package main

import (
	"abtract-factory/bakery"
	"abtract-factory/bakery/product/cake"
	"abtract-factory/bakery/product/croissant"
	"abtract-factory/bakery/product/donut"
	"fmt"
)

func main() {
	chocolateMenu, _ := bakery.GetBakeryMenu("chocolate")
	strawberryMenu, _ := bakery.GetBakeryMenu("strawberry")

	chocoDonut := chocolateMenu.MakeDonut()
	chocoCake := chocolateMenu.MakeCake()
	chocoCriosant := chocolateMenu.MakeCroissant()

	strawberryDonut := strawberryMenu.MakeDonut()
	strawberryCake := strawberryMenu.MakeCake()
	strawberryCroissant := strawberryMenu.MakeCroissant()

	fmt.Println("==========Donut Menu==========")
	printDonutDetail(chocoDonut)
	printDonutDetail(strawberryDonut)
	fmt.Println("==============================")

	fmt.Println("==========Cake Menu==========")
	printCakeDetail(chocoCake)
	printCakeDetail(strawberryCake)
	fmt.Println("==============================")

	fmt.Println("==========Croissant Menu==========")
	printCroissantDetail(chocoCriosant)
	printCroissantDetail(strawberryCroissant)
	fmt.Println("==================================")
}

func printDonutDetail(d donut.IDonut) {
	fmt.Println("donut name: ", d.GetName())
	fmt.Println("donut price: ", d.GetPrice())
}

func printCakeDetail(c cake.ICake) {
	fmt.Println("cake name: ", c.GetName())
	fmt.Println("cake price: ", c.GetPrice())
}

func printCroissantDetail(cr croissant.ICroissant) {
	fmt.Println("croissant name: ", cr.GetName())
	fmt.Println("croissant price: ", cr.GetPrice())
}
