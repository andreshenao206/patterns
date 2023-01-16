package main

import "fmt"

type Ipizza interface {
	getPrice() int
}
type tomatoTopping struct {
	pizza Ipizza
}
type cheeseTopping struct {
	pizza Ipizza
}

type vegetables struct{}

func (p *vegetables) getPrice() int {
	return 15
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {

	pizza := &vegetables{}
	fmt.Printf("Price of vegetables %d\n", pizza.getPrice())

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}
	fmt.Printf("Price of vegetables with  cheese %d\n", pizzaWithCheese.getPrice())

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of vegetables with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
