package main

import "fmt"

func main() {
	var idMap map[string]int;
	idMap = make(map[string]int);
	idMap["a"] = 1;
	idMap["b"] = 2;

	for key, val := range idMap {
		fmt.Printf("%s and value %d ",key,val);
	}
}