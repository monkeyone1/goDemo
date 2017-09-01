package main

import "fmt"

type woman struct {
	name  string
	body  string 
	face  string
	age int  
}

func main() {
	var girlfriend woman
	
	girlfriend.age = 18
	girlfriend.body = "Sex"
	girlfriend.face = "beautiful"
	girlfriend.name = "宝宝"
	fmt.Println(girlfriend)
}