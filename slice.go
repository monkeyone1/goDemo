package main

import "fmt"
func main() {
  s:=[] int {1,2,3,4,5}
  fmt.Println(s)
  fmt.Println(s[1:2],len(s),"+",cap(s),s)
  s=append(s,2)
  x:= make([]int,len(s)*2)
  copy(x,s)
  fmt.Println(x)
}