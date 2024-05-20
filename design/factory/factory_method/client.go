package factory_method

import "fmt"

func client() {
	ak47Gun, _ := getGun("ak47")
	musketGun, _ := getGun("musket")

	printDetails(ak47Gun)
	printDetails(musketGun)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
