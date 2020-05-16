package main

import "fmt"

func main(){

	var size,sum int
	
	fmt.Println("Enter the size of required array ")
	fmt.Scanln(&size)

	var arr = make ([]int,size)
	

	fmt.Println("Enter the values for array")
	for i := 0; i<size; i++{
		fmt.Scanln(&arr[i])
		sum = sum + arr[i]
		
	}
		sum= sum/len(arr)
		fmt.Println("The average is" ,sum)
}