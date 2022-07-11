package main

import "fmt"

type DonutShop struct {
	donuts    map[string]int
	customers map[string]float32
}

func (s *DonutShop) init() {
	s.donuts = map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	s.customers = make(map[string]float32)
}

func (s DonutShop) calculatePrice(count int) float32 {
	return float32(count) * 1.50
}

func (s DonutShop) placeOrder(name string, kind string, count int) {
	s.customers[name] = s.calculatePrice(count)
	s.donuts[kind] = s.donuts[kind] - count
}

func (s DonutShop) checkout(name string) {
	fmt.Printf("%s please pay %f\n", name, s.customers[name])
}

func main() {
	var name = "daryl"
	var kind = "jelly"
	var count = 5
	var donutShop = new(DonutShop)

	donutShop.init()
	donutShop.placeOrder(name, kind, count)
	donutShop.checkout(name)
	fmt.Println(donutShop.customers)
	fmt.Println(donutShop.donuts)
}
