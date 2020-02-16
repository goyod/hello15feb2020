package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pallat/hello/fizzbuzz"
	"github.com/pallat/hello/oscar"
)

func main() {
	// Hello world, the web server
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	r.HandleFunc("/oscarmale", oscarmalHandler)
	r.HandleFunc("/token", tokenHandler)

	log.Fatal(http.ListenAndServe(":8081", r))
}

func tokenHandler(w http.ResponseWriter, req *http.Request) {
	mySigningKey := []byte("GMMx4P8OI8")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		Issuer:    "pallat",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&map[string]string{
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&map[string]string{
		"token": ss,
	})
}

func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&map[string]string{
				"error": fmt.Sprintf("panic: %s", r),
			})
		}
	}()

	tokenString := req.Header.Get("Authorization")[7:]

	var getSecretKey = func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("GMMx4P8OI8"), nil
	}

	_, err := jwt.Parse(tokenString, getSecretKey)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&map[string]string{
			"error": "token is not valid",
		})
		return
	}

	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.New(n).String())
}

func oscarmalHandler(w http.ResponseWriter, req *http.Request) {
	m := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")

	w.Header().Set("Content-type", "text/json")
	json.NewEncoder(w).Encode(&m)
}
