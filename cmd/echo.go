package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
	e.GET("/fizzbuzz/:number", fizzbuzzHandler, authorizeMiddleware("GMMx4P8OI8"))
	e.POST("/fizzbuzz", postFizzBuzzHandler)

	h := &randomFizzBuzz{random: IntnFunc(rand.Intn)}
	e.GET("/fizzbuzzr", h.handler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type IntnFunc func(int) int

func (fn IntnFunc) Intn(n int) int {
	return fn(n)
}

// Handler
func oscarmale(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func authorizeMiddleware(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")[7:]

			var getSecretKey = func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(key), nil
			}

			_, err := jwt.Parse(tokenString, getSecretKey)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "token is not valid",
				})
			}

			return next(c)
		}
	}
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
