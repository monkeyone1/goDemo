package main

import "fmt"

func main() {
	a:= 100
	b:= 200
	var c = add(a, b)
	fmt.Println(c)
	x, y := swap("hello", "world")
	fmt.Println(x, y)
}

func add (num1, num2 int) int {
	
	return num1+num2
}
func swap (x,y string) (string, string) {
	return y, x+"!!!!"
}
