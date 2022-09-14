package BakeSale

import "fmt"

type ItemData struct {
	price float64
	stock int
}

var items = map[string]ItemData{}

func SetItemsValues(initItems map[string]ItemData) {
	items = initItems
}

type Sale struct {
	order        string
	currentTotal float64
	amountPaid   float64
}

type ISale interface {
	bakeSale() string
	computeChange() string
}

func (s *Sale) bakeSale() string {
	var total float64

	for i := 0; i < len(s.order); i++ {
		char := string(s.order[i])
		if char != "," {
			total += getPrice(char)
			updateStock(char)
			if items[char].stock <= 0 {
				return "Not enough stock"
			}
		}

	}
	s.currentTotal = total
	return fmt.Sprintf("$%2.2f", s.currentTotal)
}

func (s *Sale) computeChange() string {
	if s.amountPaid < s.currentTotal {
		return "Not enough money"
	}
	change := s.amountPaid - s.currentTotal
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
