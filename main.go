package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello There!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful!\n")
	a_form := r.FormValue("a")
	b_form := r.FormValue("b")
	c_form := r.FormValue("c")

	a, _ := strconv.ParseFloat(a_form, 64)
	b, _ := strconv.ParseFloat(b_form, 64)
	c, _ := strconv.ParseFloat(c_form, 64)

	var root1, root2, imaginary, discriminant float64

	discriminant = (b * b) - (4 * a * c)

	if discriminant > 0 {
		root1 = (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 = (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Fprintf(w, "Two Distinct Real Roots Exist: root1 = %v, and root2 = %v", root1, root2)
	} else if discriminant == 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		fmt.Fprintf(w, "Two Equal and Real Roots Exist: root1 = %v, and root2 = %v", root1, root2)
	} else if discriminant < 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		imaginary = math.Sqrt(-discriminant) / (2 * a)
		fmt.Fprintf(w, "Two Distinct Complex Roots Exist: root1 = %v + %v, and root2 = %v - %v", root1, imaginary, root2, imaginary)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at Port 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
