package main

import "fmt"
func main() {
	arr:= [5]int{}
	arr[0]=1;
	for i, val := range arr {
		fmt.Printf("%d and index %d",val,i);
	}
}