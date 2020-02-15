package oscar_test

import (
	"testing"

	"github.com/pallat/hello/oscar"
)

func TestActorWhoGotMoreThanOne(t *testing.T) {
	oscar.ActorWhoGotMoreThanOne("./oscar_age_male.csv")
}
