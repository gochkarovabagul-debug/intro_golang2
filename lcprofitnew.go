package main

import(
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func profit(prices []int)int{
	profit:=[]int{}
	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices); j++{
			pro:=prices[j]-prices[i]
			if j>i{
				profit=append(profit, pro)
			}
		}
	}
	bestprofit:=0
	for i:=0; i<len(profit); i++{
		if profit[i]>bestprofit{
			bestprofit=profit[i]
		}
	}
	return bestprofit
}
func profithttp(w http.ResponseWriter, r *http.Request){
	// var n int
	// fmt.Scan(&n)
	// n1:=r.FormValue("number")
	// n,_:=strconv.Atoi(n1)
	// prices:= make ([]int, n)
	// for i:=0; i<n; i++{
	// 	fmt.Scan(&prices[i])
	// }
	getnumber:=r.FormValue("numbers")
	numbers:=strings.Split(getnumber, ",")
	prices:=make([]int,0, len(numbers))
	for _, raw:=range numbers{
		v, err:=strconv.Atoi(raw)
		if err!=nil{
			fmt.Println(err)
			continue
		}
		prices=append(prices,v)
	}
	fmt.Fprintln(w, prices)
	fmt.Fprintln(w, profit(prices))
}
func main(){
	http.HandleFunc("/profit", profithttp)
	http.ListenAndServe(":8080", nil)
	// var n int
	// fmt.Scan(&n)
	// prices:= make ([]int, n)
	
	// fmt.Println(prices)
	// fmt.Println(profit(prices))
	
}