package fizzbuzz

import "strconv"


func FizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	} else if num == 5 {
		return "Buzz"
	} else if num % 3 == 0 && num % 5 == 0 {
		return "FizzBuzz"
	}
	return strconv.FormatInt(int64(num), 10)
}	