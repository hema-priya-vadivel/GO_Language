package main //main package

import "fmt" //importing format package

func main() { //main function
	var x int = 6 // default value 0
	var y int = 7
	var i, j int = 7, 8
	var a, b float32 = 1.45, 7.89
	var u float64 = 78.9
	var str string = "I love to learn GO!" //default value ""
	var bol bool = true                    //default value false
	//datatypes:int,float32,float64,string,bool
	s, t := 34, 56 //short variable declaration :=

	//constants-cannot change their value
	const pi = 3.14 // cannot be declared using shorthand declaration
	const hu = "hello this is my first go program"

	//arithmetic operators
	d := 34
	e := 2
	result := d + e
	fmt.Println(result)
	result1 := d - e
	fmt.Println(result1)
	result2 := d / e
	fmt.Println(result2)
	result3 := d % e
	fmt.Println(result3)
	result4 := d * e
	fmt.Println(result4)

	//assignment operators
	d += e
	fmt.Println(d)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(str)
	fmt.Println(bol)
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(u)
	fmt.Println(pi)
	fmt.Println(hu)

}
