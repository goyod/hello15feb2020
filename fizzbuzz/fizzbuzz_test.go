package fizzbuzz

import "testing"

func TestFizzzBuzzGivenOneSayOne(t *testing.T) {
	var given = 1
	var want = "1"

	get := Say(given)
	if want != get {
		t.Errorf("given %v want %q but get %q", given, want, get)
	}
}
