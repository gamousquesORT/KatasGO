package fizzbuzz

import ("testing")

/*
Write a program that prints one line for each number from 1 to 100
For multiples of three print Fizz instead of the number
For the multiples of five print Buzz instead of the number
For numbers which are multiples of both three and five print FizzBuzz instead of the number

*/

func TestFizzBuzz(t *testing.T) {
	t.Run("Should return Fizz for 3", func(t *testing.T) {
		got := FizzBuzz(3)
		want := "Fizz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Should return Buzz for 5", func(t *testing.T) {
		got := FizzBuzz(5)
		want := "Buzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Should return FizzBuzz for 15", func(t *testing.T) {
		got := FizzBuzz(15)
		want := "FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Should return 1 for 1", func(t *testing.T) {
		got := FizzBuzz(1)
		want := "1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Should return 15 for FizzBuzz", func(t *testing.T) {
		got := FizzBuzz(15)
		want := "FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Should return 10 for 10", func(t *testing.T) {
		got := FizzBuzz(10)
		want := "10"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}