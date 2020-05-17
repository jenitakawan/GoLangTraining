package main

import "fmt"

func main(){
	slice:=[]int{1,2,3}

	for i :=len(slice)/2-1;i>=0;i--{
		reverse :=len(slice)-1-i
		slice[i],slice[reverse]=slice[reverse],slice[i]
	}
	fmt.Println(slice)


}