// package main

// import "fmt"

// func tersi(x string) string {
// 	if len(x) <= 1 {
// 		return x
// 	} else {
// 		return tersi(x[1:]) + string(x[0])
// 	}
// }
// func main() {
// 	var x string
// 	fmt.Printf("soz giriz: \n")
// 	fmt.Scan(&x)
// 	fmt.Printf("sozun tersi: %v\n", tersi(x))
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b int
// 	for a := 0; a <= b; a++ {
// 		if a-b == 0 {
// 			fmt.Print(a)
// 		}
// 	}

//		for b := 0; b <= a; b++ {
//			if a-b == 0 {
//				fmt.Print(b)
//			}
//		}
//		fmt.Printf("san: %v", a)
//		fmt.Scan(&a)
//		fmt.Print(b)
//	}
package main

import "fmt"

type Animalfam struct {
	animal []animal
}

// test git

type animal struct {
	name     string
	weight   int
	consumes int
	produces int
}

func totalconsumes(sliceofanimal []animal) int {
	var consumes int
	for _, a := range sliceofanimal {
		consumes += a.consumes
	}
	return consumes
}
func totalproduces(sliceofanimal []animal) int {
	var produces int
	for _, a := range sliceofanimal {
		produces += a.produces
	}
	return produces
}
func totalprofit(sliceofanimal []animal) int {

	return totalproduces(sliceofanimal) - totalconsumes(sliceofanimal)

}
func totalweight(sliceofanimal []animal) int {
	var weight int
	for _, a := range sliceofanimal {
		weight += a.weight
	}
	return weight
}

func main() {

	sliceofanimal := []animal{
		{name: "camel", weight: 40, consumes: 3, produces: 4},
		{name: "cow", weight: 30, consumes: 2, produces: 5},
		{name: "dog", weight: 20, consumes: 1, produces: 6},
	}
	fmt.Println("total comsumes: ", totalconsumes(sliceofanimal))
	fmt.Println("total producees: ", totalproduces(sliceofanimal))
	fmt.Println("total profit: ", totalprofit(sliceofanimal))
	fmt.Println("total weight: ", totalweight(sliceofanimal))

	sameweightanimals := []string{}
	for _, a := range sliceofanimal {
		if a.weight == 40 {
			sameweightanimals = append(sameweightanimals, a.name)
		}
	}
	fmt.Println("same weight animals:", sameweightanimals)
}

// func totalprofit(sliceofanimal []animal) int{
//  profit:= func (totalproduces(sliceofanimal []animal))-func (totalconsumes(sliceofanimal []animal))
//  return profit
// }
