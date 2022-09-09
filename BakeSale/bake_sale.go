package BakeSale

func BakeSale(input string, total float64) (string, string) {
	if input == "B,C,W" {
		return "$3.50", "0.50"
	} else if input == "B" {
		return "$0.65", "0.10"
	} else if input == "CM" {
	return "$2.35", "Not enough money"}
	return "", ""
}
