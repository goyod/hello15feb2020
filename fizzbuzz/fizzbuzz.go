package fizzbuzz

import "strconv"

type FizzBuzz int

func New(n int) FizzBuzz {
	return FizzBuzz(n)
}

func (f FizzBuzz) String() string {
	return Say(int(f))
}

func Say(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
