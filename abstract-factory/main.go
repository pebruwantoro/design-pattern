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
	blueberryMenu, _ := bakery.GetBakeryMenu("blueberry")

	chocoDonut := chocolateMenu.MakeDonut()
	chocoCake := chocolateMenu.MakeCake()
	chocoCroissant := chocolateMenu.MakeCroissant()
	// chocoCookies := chocolateMenu.MakeCookies()

	strawberryDonut := strawberryMenu.MakeDonut()
	strawberryCake := strawberryMenu.MakeCake()
	strawberryCroissant := strawberryMenu.MakeCroissant()
	// strawberryCookies := strawberryMenu.MakeCookies()

	blueberryDonut := blueberryMenu.MakeDonut()
	blueberryCake := blueberryMenu.MakeCake()
	blueberryCroissant := blueberryMenu.MakeCroissant()
	// blueberryCookies := blueberryMenu.MakeCookies()

	fmt.Println("==========Donut Menu==========")
	printDonutDetail(chocoDonut)
	printDonutDetail(strawberryDonut)
	printDonutDetail(blueberryDonut)
	fmt.Println("==============================")

	fmt.Println("==========Cake Menu==========")
	printCakeDetail(chocoCake)
	printCakeDetail(strawberryCake)
	printCakeDetail(blueberryCake)
	fmt.Println("==============================")

	fmt.Println("==========Croissant Menu==========")
	printCroissantDetail(chocoCroissant)
	printCroissantDetail(strawberryCroissant)
	printCroissantDetail(blueberryCroissant)
	fmt.Println("==================================")

	// fmt.Println("==========Cookies Menu==========")
	// printCookiesDetail(chocoCookies)
	// printCookiesDetail(strawberryCookies)
	// printCookiesDetail(blueberryCookies)
	// fmt.Println("==================================")
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

// func printCookiesDetail(cr cookies.ICookies) {
// 	fmt.Println("cookies name: ", cr.GetName())
// 	fmt.Println("cookies price: ", cr.GetPrice())
// }
