package fizzbuzz

import "fmt"

func Fizzbuzz(number int) string {
	switch 0 {
	case number % 15:
		return "FizzBuzz"
	case number % 5:
		return "Buzz"
	case number % 3:
		return "Fizz"
	default:
		return fmt.Sprintf("%d", number)
	}
}
