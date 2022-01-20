package decorator

import (
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	mypizza := &chickenPizza{}
	mypizzaWithTomato := NewTomatoTopping(mypizza)
	mypizzaWithTomatoAndCheese := NewCheeseTopping(mypizzaWithTomato)
	fmt.Printf("Price of chicken pizza with tomato and cheese topping is %d\n", mypizzaWithTomatoAndCheese.getPrice())
}
