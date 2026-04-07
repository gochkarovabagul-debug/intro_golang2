// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// func check(n int) bool {
// 	if n%3 == 0 {
// 		return true
// 	}
// 	return false
// }

// func first(w http.ResponseWriter, r *http.Request) {
// 	number := r.FormValue("san")
// 	number2, _ := strconv.Atoi(number)
// 	if check(number2) {
// 		fmt.Fprint(w, "san uce bolunya")
// 	} else {
// 		fmt.Fprint(w, "san uce bolunenok")
// 	}

// }

//	func main() {
//		http.HandleFunc("/check", first)
//		http.ListenAndServe(":8000", nil)
//	}
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
func howru(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "good thanks, what about you")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/mood", howru)
	go http.ListenAndServe(":8080", nil)
	go http.ListenAndServe(":7777", nil)
	http.ListenAndServe(":8000", nil)
}

// }	mux := http.NewServeMux()
// http.ListenAndServe(":8000", mux)
// func main() {
// 	min := http.NewServeMux() // 🔥 öz router

// 	min.HandleFunc("/hello", hello)
// 	// mux.HandleFunc("/mood", howru)

// 	http.ListenAndServe(":8000", min) // 🔥 nil ýerine mux
// }
