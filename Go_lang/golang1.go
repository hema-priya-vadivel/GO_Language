package main

import "fmt"

func main() {
	//Relational Operators
	a := 3
	b := 7
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	//Logical operators
	d := 34
	e := 45
	fmt.Println((d < e) && (e > d))
	fmt.Println((d < e) || (e < d))
	fmt.Println(!(d > e))
}
