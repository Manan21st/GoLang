package main

import "fmt"

type Person struct {
	name string
	age int
	phone string
}
func main() {
	var p1 Person;
	p1.name = "John";
	p1.age = 25;
	p1.phone = "1234567890";

	p2 := new(Person);
	fmt.Println(p1);
	fmt.Println(p2);
}