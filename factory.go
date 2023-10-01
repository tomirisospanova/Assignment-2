package main

import "fmt"

type GameObjectFactory interface {
	Create() GameObject
}

type GameObject interface {
	Display()
}

type Ninja struct{}
type Enemy struct{}
type Item struct{}

func (ninja *Ninja) Display() {
	fmt.Println("Ниндзя атакует!")
}

func (enemy *Enemy) Display() {
	fmt.Println("Враг атакует!")
}

func (item *Item) Display() {
	fmt.Println("Подобран предмет!")
}

// Фабрика для создания ниндзя
type NinjaFactory struct{}

// Фабрика для создания врагов
type EnemyFactory struct{}

// Фабрика для создания предметов
type ItemFactory struct{}

func (factory *NinjaFactory) Create() GameObject {
	return &Ninja{}
}

func (factory *EnemyFactory) Create() GameObject {
	return &Enemy{}
}

func (factory *ItemFactory) Create() GameObject {
	return &Item{}
}

func main() {
	ninjaFactory := &NinjaFactory{}
	enemyFactory := &EnemyFactory{}
	itemFactory := &ItemFactory{}

	ninja := ninjaFactory.Create()
	enemy := enemyFactory.Create()
	item := itemFactory.Create()

	ninja.Display()
	enemy.Display()
	item.Display()
}
