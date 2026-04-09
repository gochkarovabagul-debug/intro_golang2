package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func Even(n2 int, m int) int {

	if n2 == 1 {
		return m
	}
	for {
		m += 1
		if isEven(m) {
			return Even(n2-1, m)
		}
	}
}
func even(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("number")
	n2, _ := strconv.Atoi(n)
	fmt.Fprintln(w, Even(n2, 2))
}
func main() {
	http.HandleFunc("/even", even)
	http.ListenAndServe(":8000", nil)
}
