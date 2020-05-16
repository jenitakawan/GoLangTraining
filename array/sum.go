package main

import "fmt"

func main(){
	array:=[6]int{1,4,5,68,88,90}
	sum:=0
	for i:=0;i<len(array);i++{
		sum = sum + array[i]
		
	}
	fmt.Println("The sum is",sum)
}