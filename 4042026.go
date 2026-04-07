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
func barlama(w http.ResponseWriter, r *http.Request) {

	n2 := r.FormValue("number")
	n1, _ := strconv.Atoi(n2)
	if isPrime(n1) {
		fmt.Fprintln(w, "yonekey")
	} else {
		fmt.Fprintln(w, "yonekey dal")
	}
	fmt.Fprintln(w, "barlamak un basgada san girizip bilersiniz")

}
func salamlasmak(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "salam, hos geldiniz")
}

func main() {
	http.HandleFunc("/prime", salamlasmak)
	http.HandleFunc("/hello", barlama)
	http.ListenAndServe(":8000", nil)
}
