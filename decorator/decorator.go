package decorator

type pizza interface {
	getPrice() int
}

type chickenPizza struct{}

func (*chickenPizza) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza pizza
}

func NewTomatoTopping(pizza pizza) *tomatoTopping {
	return &tomatoTopping{
		pizza: pizza,
	}
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 1
}

type cheeseTopping struct {
	pizza pizza
}

func NewCheeseTopping(pizza pizza) *cheeseTopping {
	return &cheeseTopping{
		pizza: pizza,
	}
}

func (t *cheeseTopping) getPrice() int {
	return t.pizza.getPrice() + 2
}
