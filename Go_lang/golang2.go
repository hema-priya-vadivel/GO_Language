package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter the string")
	fmt.Scanln(&str)
	fmt.Println("The string you entered is " + str)

	var hello int
	fmt.Println("Enter the integer value")
	fmt.Scanln(&hello)
	fmt.Println(hello)

	if mark := 67; mark > 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	res := 0
	for i := 0; i < 10; i += 3 {
		res += i
	}
	fmt.Println(res)
}
