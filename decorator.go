package main

import "fmt"

type CoffeeComponent interface {
	Cost() int
	Description() string
}
//Еspresso
type Espresso struct{}

func (e *Espresso) Cost() int {
	return 100
}

func (e *Espresso) Description() string {
	return "Espresso"
}
//latte
type Latte struct{}

func (l *Latte) Cost() int {
	return 120
}

func (l *Latte) Description() string {
	return "Latte"
}

// Молоко
type MilkDecorator struct {
	component CoffeeComponent
}

func (m *MilkDecorator) Cost() int {
	return m.component.Cost() + 20
}

func (m *MilkDecorator) Description() string {
	return m.component.Description() + ", with Milk"
}

// Сахар
type SugarDecorator struct {
	component CoffeeComponent
}

func (s *SugarDecorator) Cost() int {
	return s.component.Cost() + 10
}

func (s *SugarDecorator) Description() string {
	return s.component.Description() + ", with Sugar"
}

func main() {
	// Espresso
	espresso := &Espresso{}
	fmt.Printf("Cost: %d, Description: %s\n", espresso.Cost(), espresso.Description())

	// Добавляем молоко 
	espressoWithMilk := &MilkDecorator{espresso}
	fmt.Printf("Cost: %d, Description: %s\n", espressoWithMilk.Cost(), espressoWithMilk.Description())

	// Создаем Latte
	latte := &Latte{}
	fmt.Printf("Cost: %d, Description: %s\n", latte.Cost(), latte.Description())

	// Сахар к Latte
	latteWithSugar := &SugarDecorator{latte}
	fmt.Printf("Cost: %d, Description: %s\n", latteWithSugar.Cost(), latteWithSugar.Description())
}
