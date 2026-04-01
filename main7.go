package main

import "fmt"

func kwadrat(x,y int) int{
	
	var kwadrat=x; 
	for i:=1; i <= y; i++{
		kwadrat*=kwadrat
		return kwadrat
	} 
}
