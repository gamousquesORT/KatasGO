package BakeSale

import "fmt"

type Item struct {
	price float64
	stock int
}

// var items = make(map[string]Item)
var items = map[string]Item{"B": {0.65, 48}, "M": {1.0, 36}, "C": {1.35, 24}, "W": {1.5, 2}}

func BakeSale(input string) string {
	var total float64
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char != "," {
			total += getPrice(char)
			updateStock(char)
			if items[char].stock <= 0 {
				return "Not enough stock"
			}
		}

	}

	return fmt.Sprintf("$%2.2f", total)
}

func updateStock(char string) {
	actItem := items[char]
	actItem.stock--
	items[char] = actItem
}

func getPrice(prod string) float64 {
	return items[prod].price
}
