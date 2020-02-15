package fizzbuzz

import "testing"

func TestFizzBuzzSayOrigin(t *testing.T) {
	t.Run("given 1 say 1", func(t *testing.T) {
		var given = 1
		var want = "1"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 2 say 2", func(t *testing.T) {
		var given = 2
		var want = "2"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 4 say 4", func(t *testing.T) {
		var given = 4
		var want = "4"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 7 say 7", func(t *testing.T) {
		var given = 7
		var want = "7"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}

func TestFizzBuzzSayFizz(t *testing.T) {
	t.Run("given 3 say Fizz", func(*testing.T) {
		var given = 3
		var want = "Fizz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 6 say Fizz", func(t *testing.T) {
		var given = 6
		var want = "Fizz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 9 say Fizz", func(t *testing.T) {
		var given = 9
		var want = "Fizz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}

func TestFizzBuzzSayBuzz(t *testing.T) {
	t.Run("given 5 say Buzz", func(t *testing.T) {
		var given = 5
		var want = "Buzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 10 say Buzz", func(t *testing.T) {
		var given = 10
		var want = "Buzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 20 say Buzz", func(t *testing.T) {
		var given = 20
		var want = "Buzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}

func TestFizzBuzzSayFizzBuzz(t *testing.T) {
	t.Run("given 15 say FizzBuzz", func(t *testing.T) {
		var given = 15
		var want = "FizzBuzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 30 say FizzBuzz", func(t *testing.T) {
		var given = 30
		var want = "FizzBuzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 45 say FizzBuzz", func(t *testing.T) {
		var given = 45
		var want = "FizzBuzz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}
