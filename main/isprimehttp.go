package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func nextPrime(n int, min int) int {
	if n == 1 {
		return min
	}
	for {
		min += 1
		if isPrime(min) {
			return nextPrime(n-1, min)
		}
	}
}
func prime(w http.ResponseWriter, r *http.Request) {
	n1 := r.FormValue("number")
	n, _ := strconv.Atoi(n1)
	fmt.Fprintln(w, nextPrime(n, 2))
}
func main() {
	http.HandleFunc("/prime", prime)
	http.ListenAndServe(":8080", nil)
}
