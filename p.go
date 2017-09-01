// 指针
package main

import "fmt"
func main() {
	const Max int = 3
	a:= []int{10,100,200}
	var ptr [Max]*int
	for i := 0; i < Max; i++ {
		ptr[i] = &a[i]
	}
	for i := 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d ",i,*ptr[i])
	}
	for i := 0; i < Max; i++ {
		fmt.Println(a[i])
	}
}


