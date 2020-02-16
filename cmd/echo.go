package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pallat/hello/fizzbuzz"
	"github.com/pallat/hello/oscar"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/oscarmale", oscarmale)
	e.GET("/fizzbuzz/:number", fizzbuzzHandler)
	e.POST("/fizzbuzz", postFizzBuzzHandler)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	h := &randomFizzBuzz{random: r1}
	e.GET("/fizzbuzzr", h.handler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func oscarmale(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func fizzbuzzHandler(c echo.Context) error {
	numberString := c.Param("number")
	n, _ := strconv.Atoi(numberString)
	return c.String(http.StatusOK, fizzbuzz.Say(n))
}

type randomer interface {
	Intn(int) int
}

type randomFizzBuzz struct {
	random randomer
}

func (r *randomFizzBuzz) handler(c echo.Context) error {
	n := r.random.Intn(100)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"number":  n,
		"message": fizzbuzz.Say(n),
	})
}

func randomFizzBuzzHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, fizzbuzzController(rand.Intn(100)))
}

func fizzbuzzController(n int) map[string]interface{} {
	return map[string]interface{}{
		"number":  n,
		"message": fizzbuzz.Say(n),
	}
}

func postFizzBuzzHandler(c echo.Context) error {
	var req map[string]int
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	type fizzbuzzResponse struct {
		Number   int    `json:"number"`
		FizzBuzz string `json:"fizzbuzz"`
	}

	return c.JSON(http.StatusOK, fizzbuzzResponse{
		Number:   req["number"],
		FizzBuzz: fizzbuzz.Say(req["number"]),
	})
}
