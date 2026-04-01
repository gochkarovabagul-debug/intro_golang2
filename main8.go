// package main

// import (
// 	"fmt"
// )

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 11; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }

// package main

// import (
// 	"fmt"
// )

// func Sqrt(x float64) float64 {
// 	z := 1.0

// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}

// 	return z
// }

//	func main() {
//		fmt.Println(Sqrt(2))
//	}
// package main

// import "fmt"

// func main() {
// 	var num int

// 	fmt.Print("Enter a number : ")
// 	fmt.Scan(&num)

// 	switch num {
// 	case 1:
// 		fmt.Println("You chose ONE")
// 	case 2:
// 		fmt.Println("You chose TWO")
// 	case 3:
// 		fmt.Println("You chose THREE")
// 	default:
// 		fmt.Println("Invalid number")
// 	}
// }

// package main

// import "fmt"

// func reverse(s string) string {
// 	// base case
// 	if len(s) <= 1 {
// 		return s
// 	}

// 	// recursive step
// 	return reverse(s[1:]) + string(s[0])
// }

// func main() {
// 	word := "hello"
// 	fmt.Println(reverse(word))
// }
////////////////////////////////////????????????????//////////////////////////////////////////////////////////////
package main

import "fmt"

func tersi(x string) string {
	if len(x) <= 1 {
		return x
	} else {
		return tersi(x[1:]) + string(x[0])
	}
}
func main() {
	var x string
	fmt.Printf("soz giriz: \n")
	fmt.Scan(&x)
	fmt.Printf("sozun tersi: %v\n", tersi(x))
}
// ///////////////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

//	func main() {
//		var jogap string
//		fmt.Printf("sen akyllymy:")
//		fmt.Scan(&jogap)
//		p := &jogap
//		*p = "akmak"
//		if jogap=*p{
//			return malades
//		}
//		return *p
//	}
// package main

// import "fmt"

// func main() {
// 	var jogap string

// 	fmt.Print("sen akyllymy: ")
// 	fmt.Scan(&jogap)

// 	p := &jogap

// 	if *p == "yes" {
// 		fmt.Println("malades")
// 	} else {
// 		fmt.Println(*p)
// 	}
// }

// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	fmt.Println("old point was: ", v.Y)
// 	v.Y = 1
// 	fmt.Println("new point is:", v.Y)
// }



// package main

// import "fmt"

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)

	
// 	s = s[:0]
// 	printSlice(s)

	
// 	s = s[:4]
// 	printSlice(s)

// 	s = s[2:]
// 	printSlice(s)
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
