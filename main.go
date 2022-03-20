package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type myHandler int

func getNumber(queryParam string, r *http.Request) int {
	numberSlice := r.URL.Query()[queryParam]
	var number int = 0
	if len(numberSlice) > 0 {
		number, _ = strconv.Atoi(numberSlice[0])
	}
	return number
}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	number1 := getNumber("number1", r)
	number2 := getNumber("number2", r)
	added := strconv.Itoa(number1 + number2)
	fmt.Fprintln(w, added)
}

func main() {
	var handler myHandler
	log.Fatal(http.ListenAndServe(":8080", handler))
}
