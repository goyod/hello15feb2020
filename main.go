package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pallat/hello/fizzbuzz"
	"github.com/pallat/hello/oscar"
)

func main() {
	// Hello world, the web server
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	r.HandleFunc("/oscarmale", oscarmalHandler)

	log.Fatal(http.ListenAndServe(":8081", r))
}

func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.Say(n))
}

func oscarmalHandler(w http.ResponseWriter, req *http.Request) {
	m := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")

	w.Header().Set("Content-type", "text/json")
	json.NewEncoder(w).Encode(&m)
}
