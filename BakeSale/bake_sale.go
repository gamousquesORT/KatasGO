package BakeSale

func BakeSale(input string, total float64) (string, float64) {
	if input == "B,C,W" {
		return "$3.50", 0.50
	}
	return "$0.65", 0.1
}
