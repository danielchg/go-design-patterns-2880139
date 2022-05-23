package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n--------------------")

	// TODO: make an americano coffee using the coffeemachine API
	americano := CoffeeMachine{}
	// determine beans amount to use - 5oz for every 8oz size
	beansAmount := (size / 8) * 5.0
	americano.startCoffee(beansAmount, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is ready!")

}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n--------------------")

	// TODO: make a latte coffee using the coffeemachine API
	latte := &CoffeeMachine{}

	// determine beans amount to use - 5oz for every 8oz size
	beansAmount := (size / 8) * 5.0
	// determine milk amount to use - 2oz for every 8oz size
	milkAmount := (size / 8) * 2.0

	latte.startCoffee(beansAmount, 2)
	latte.grindBeans()
	latte.useMilk(milkAmount)
	latte.doFoam(foam)
	latte.endCoffee()

	fmt.Println("Latte is ready!")
}
