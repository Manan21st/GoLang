package main

import ("fmt"
	"sort" 
	"strconv")

func main(){
	slice := make([]int, 0, 3)
	
	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		var input string
		fmt.Scan(&input)
		
		if input == "X" {
			break
		}
		
		num,err := strconv.Atoi(input)
		if err != nil {}
		
		slice = append(slice, num)
		
		sort.Ints(slice)
		fmt.Println("Sorted slice:", slice)
	}
	
	fmt.Println("Program exited.")
}