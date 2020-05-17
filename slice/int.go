package main

import (
	 "fmt"
	 "strconv"
	  //"Ints"
	  "sort"	
	//  "log"

)

//import "sort"



func main(){

	
	var arr = make ([]string,5)
	fmt.Println("Enter the values for array")
	
	for i := 0; i<5; i++{
		fmt.Scanln(&arr[i])
	
	}

	var vals[]int

	for i := range arr{
		num :=arr[i]
		num1,err:= strconv.Atoi(num)
		
		if err ==nil{
			vals = append(vals,num1)
		}
	}
		
			fmt.Println("The numbers are",vals)
			sort.Ints(vals)
			fmt.Println("The sorted numbers are",vals)

		
}