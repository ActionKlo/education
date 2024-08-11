package main //nolint:typecheck
import "fmt"

type Size int

const (
	small Size = iota
	medium
	big
)

type Pizza struct {
	size      Size
	cheese    bool
	pepperoni bool
	mushrooms bool
	olives    bool
}

type pizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() PizzaBuilder {
	return &pizzaBuilder{&Pizza{}}
}

type PizzaBuilder interface {
	SetSize(size Size) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	AddOlives() PizzaBuilder
	Build() *Pizza
}

func (b *pizzaBuilder) SetSize(size Size) PizzaBuilder {
	b.pizza.size = size
	return b
}

func (b *pizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.cheese = true
	return b
}

func (b *pizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.pepperoni = true
	return b
}

func (b *pizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.mushrooms = true
	return b
}

func (b *pizzaBuilder) AddOlives() PizzaBuilder {
	b.pizza.olives = true
	return b
}

func (b *pizzaBuilder) Build() *Pizza {
	return b.pizza
}

func main() {
	smallPepperoni := NewPizzaBuilder().SetSize(small).AddCheese().AddPepperoni().Build()
	mediumFull := NewPizzaBuilder().SetSize(medium).AddCheese().AddPepperoni().AddMushrooms().AddOlives().Build()

	fmt.Println("small pepperoni:", smallPepperoni)
	fmt.Println("medium full:", mediumFull)
}
