package BakeSale

import "fmt"

type Item struct {
	price float64
	stock int
}

type Sales interface {
	bakeSale() string
}

// var items = make(map[string]Item)
var items = map[string]Item{}

func SetItemsValues(initItems map[string]Item) {
	items = initItems
}

var currentTotal = 0.0

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
	currentTotal = total

	return fmt.Sprintf("$%2.2f", total)
}

func ComputeChange(amountPaid float64) string {
	if amountPaid < currentTotal {
		return "Not enough money"
	}
	change := amountPaid - currentTotal
	return fmt.Sprintf("$%2.2f", change)
}

func updateStock(char string) {
	actItem := items[char]
	actItem.stock--
	items[char] = actItem
}

func getPrice(prod string) float64 {
	return items[prod].price
}
