package main

import (
	"fmt"
	"github.com/g3offrey/fizzbuzz/fizzbuzz"
	"log"
)

func main() {
	err := fizzbuzz.FizzBuzz(50, func(result string) {
		fmt.Println(result)
	})

	if err != nil {
		log.Fatalln(err)
	}
}
