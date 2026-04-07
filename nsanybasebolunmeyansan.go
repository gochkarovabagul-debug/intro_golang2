// package main

// import "fmt"

//	func check(n int) bool {
//		if n%5 == 0 {
//			return false
//		}
//		return true
//	}
//
//	func main() {
//		var N int
//		fmt.Scan(&N)
//		slice := []int{}
//		n := 1
//		for len(slice) < N {
//			if check(n) {
//				slice = append(slice, n)
//			}
//			n++
//		}
//		fmt.Println(slice)
//	}
package main

import "fmt"

func check(n int) bool {
	if n%5 != 0 {
		return true
	}
	return false
}
func main() {
	var N int
	fmt.Scan(&N)
	slice := []int{}
	n := 1
	for len(slice) < N {
		if check(n) {
			slice = append(slice, n)
		}
		n++
	}
	fmt.Println(slice)
}
