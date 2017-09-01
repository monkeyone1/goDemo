package main 

import "fmt"

func main() {
	a:= 10
	if a<= 10 {
		fmt.Printf("a 小于10 \n")
	} else {
		fmt.Printf("a的值是:%d",a)
	}
	i:=0
	for true {
		fmt.Println("hello")
		i++
		if i > 5 {
			break;
		}
		
	}
}