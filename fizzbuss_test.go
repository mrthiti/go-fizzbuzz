package fizzbuzz // --> (1)

import (
	"fmt"
	"testing" // --> (2)
)

func TestInputDivideBy15ShouldBeDisplayFizzBuzz(t *testing.T) { // --> (3)

	inputList := []int{15, 30, 45, 60, 75, 90}

	for _, n := range inputList {
		v := Fizzbuzz(n)

		if "FizzBuzz" != v {
			t.Error("fizzbuzz of", n, "should be 'FizzBuzz' but have", v)
		}
	}
}

func TestInputDivideBy5ShouldBeDisplayBuzz(t *testing.T) { // --> (4)

	inputList := []int{5, 10, 20, 25, 35, 40, 50}

	for _, n := range inputList {
		v := Fizzbuzz(n)

		if "Buzz" != v {
			t.Error("fizzbuzz of", n, "should be 'Buzz' but have", v)
		}
	}
}

func TestInputDivideBy3ShouldBeDisplayFizz(t *testing.T) { // --> (5)

	inputList := []int{3, 6, 9, 12, 18, 21, 24}

	for _, n := range inputList {
		v := Fizzbuzz(n)

		if "Fizz" != v {
			t.Error("fizzbuzz of", n, "should be 'Fizz' but have", v)
		}
	}
}

func TestInputOtherNumberShouldBeDisplayInputNumber(t *testing.T) { // --> (6)

	inputList := []int{1, 2, 4, 7, 11, 13, 17, 16}

	for _, n := range inputList {
		v := Fizzbuzz(n)

		if fmt.Sprintf("%d", n) != v {
			t.Error("fizzbuzz of", n, "should be '", n, "' but have", v)
		}
	}
}
