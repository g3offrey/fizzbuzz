package fizzbuzz

import (
	"errors"
	"strconv"
)

func isDivisibleByX(number int, x int) (bool, error) {
	if x == 0 {
		return false, errors.New("can't divide by 0")
	}

	return number%x == 0, nil
}

type OutputHandler func(string)

func FizzBuzz(maxValue int, handler OutputHandler) error {
	for i := 1; i <= maxValue; i++ {
		isDivisibleBy3, err := isDivisibleByX(i, 3)
		if err != nil {
			return err
		}

		isDivisibleBy5, err := isDivisibleByX(i, 5)
		if err != nil {
			return err
		}

		result := ""
		switch true {
		case isDivisibleBy3 && isDivisibleBy5:
			result = "fizzbuzz"
		case isDivisibleBy3:
			result = "fizz"
		case isDivisibleBy5:
			result = "buzz"
		default:
			result = strconv.Itoa(i)
		}

		handler(result)
	}

	return nil
}
