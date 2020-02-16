package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

type fakeIntn struct{}

func (fakeIntn) Intn(int) int {
	return 15
}

type spyEchoContext struct {
	statusCode int
	payload    interface{}
	echo.Context
}

func (m *spyEchoContext) JSON(code int, payload interface{}) error {
	m.statusCode = code
	m.payload = payload
	return nil
}

func TestRandomFizzBuzzHandler(t *testing.T) {
	h := randomFizzBuzz{random: fakeIntn{}}
	c := &spyEchoContext{}
	h.handler(c)

	if c.statusCode != http.StatusOK {
		t.Error("")
	}

	want := map[string]interface{}{
		"message": "FizzBuzz",
		"number":  15,
	}

	if v, ok := c.payload.(map[string]interface{})["message"]; ok {
		if v != "FizzBuzz" {
			t.Error("")
		}
	}

	if !reflect.DeepEqual(want, c.payload) {
		t.Error("")
	}
}
