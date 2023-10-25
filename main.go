package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2 + 2 is %d", sum))
}

// addValues adds 2 integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValues(100, 0)

	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")

		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func DivideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}

	return x / y, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
