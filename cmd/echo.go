package main

import (
	"encoding/xml"
	"net/http"
	"strconv"

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

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Soap    string   `xml:"soap,attr"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Add Add `xml:"Add"`
}

type Add struct {
	Xmlns string `xml:"xmlns,attr"`
	IntA  string `xml:"intA"`
	IntB  string `xml:"intB"`
}
