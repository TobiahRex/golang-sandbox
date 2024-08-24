package decorator_iface

import "fmt"

func main() {
	coffee := SimpleCoffee{}
	coffeeDesc := coffee.GetDescription()

	milkCoffee := MilkDecorator{Coffee: coffee}
	milkCoffeeDesc := milkCoffee.GetDescription()

	sugarMilkCoffee := SugarMilkCoffeeDecorator{Coffee: milkCoffee}
	sugarMilkCoffeeDesc := sugarMilkCoffee.GetDescription()

	fmt.Println(
		"Coffee: ", coffeeDesc,
		"\nMilk Coffee: ", milkCoffeeDesc,
		"\nSugar+Milk Coffee: ", sugarMilkCoffeeDesc)
}

type Coffee interface {
	GetDescription() string
	GetCost() float64
}

type SimpleCoffee struct {
	Cost        float64
	Description string
}

func (sc SimpleCoffee) GetDescription() string {
	return sc.Description
}

func (sc SimpleCoffee) GetCost() float64 {
	return sc.Cost
}

type MilkDecorator struct {
	Coffee Coffee
}

func (md MilkDecorator) GetDescription() string {
	return md.Coffee.GetDescription() + ", milk"
}

func (md MilkDecorator) GetCost() float64 {
	return md.Coffee.GetCost() + 0.50
}

type SugarMilkCoffeeDecorator struct {
	Coffee Coffee
}

func (smc SugarMilkCoffeeDecorator) GetDescription() string {
	return smc.Coffee.GetDescription() + ", sugar"
}

func (smc SugarMilkCoffeeDecorator) GetCost() float64 {
	return smc.GetCost() + 0.25
}
