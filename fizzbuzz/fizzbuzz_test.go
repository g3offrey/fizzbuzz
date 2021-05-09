package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	maxValue := 15
	handlersCalls := []string{}
	fakeHandler := func(result string) {
		handlersCalls = append(handlersCalls, result)
	}

	FizzBuzz(maxValue, fakeHandler)

	if len(handlersCalls) != maxValue {
		t.Errorf("OutputHandler should have been called %d times but was called %d times", maxValue, len(handlersCalls))
	}

	expectedHandlersCalls := []string{
		"1", "2", "fizz", "4", "buzz",
		"fizz", "7", "8", "fizz", "buzz",
		"11", "fizz", "13", "14", "fizzbuzz",
	}
	callsAreEqual := reflect.DeepEqual(handlersCalls, expectedHandlersCalls)
	if !callsAreEqual {
		t.Errorf("Actual calls %v != Expected calls %v", handlersCalls, expectedHandlersCalls)
	}

}
