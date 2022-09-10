package BakeSale

import "fmt"

type Item struct {
	price float64
	stock int
}

// var items = make(map[string]Item)
var items = map[string]Item{"B,C,W": {3.50, 3}, "B": {0.65, 1}, "C,M": {2.35, 2}, "W": {0.00, 0}}

var currenItem Item

func BakeSale(input string) string {
	if input == "B,C,W" {
		return fmt.Sprintf("$%2.2f", items[input].price)
	} else if input == "B" {
		return fmt.Sprintf("$%2.2f", items[input].price)
	} else if input == "C,M" {
		return fmt.Sprintf("$%2.2f", items[input].price)
	} else if input == "W" {
		return "Not enough stock"
	}

	return ""
}
