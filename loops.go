package main

import (
	"fmt"
	"bufio"
	"os"
)

func readLine(greeting string) string{
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _:= reader.ReadLine()
	return string(text)
}


func main() {
	for i:=1; i < 10; i+=2{
		fmt.Println(i)
	}
	var country = readLine("Where are you from: ")
	if country == "Kyrgyzstan" || country == "Kazakhstan" || country == "Uzbekistan" || country == "Tajikistan" || country == "Turkmenistan"{
		fmt.Println("You from Central Asia")
	}else{
		fmt.Println("You from other country")
	}
	fmt.Println("Goodbye")
}