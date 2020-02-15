// +build integration

package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/pallat/hello/fizzbuzz"
)

func TestFizzBuzz1To100(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Print(fizzbuzz.Say(i) + ",")
	}
}
