package main

import (
	"fmt" 
)

func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
   }  

func averageNumbers(numbers[]int) float32{
	return float32(sum(numbers)) / float32(len(numbers))
}

func hello(){
	fmt.Println("hello world, bye")
}

func main() {
	fmt.Printf("Avarage numbers [10,4,6,5,7,5]: %.1f", averageNumbers([]int{10,4,6,5,7,5,5,6,5}))
	hello()

}


