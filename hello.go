package main

import "fmt"

var (
    a int = 23
    b bool = false
)
const PI =3.14

func main() {
    fmt.Println("Hello, 世界/n")
    var age int = 123
    var name string = "hello"
    // , 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
    fmt.Println(age,name)
    your_name := "zzd"
    fmt.Println(your_name)
    fmt.Printf("面积为 : %f", PI)

    

}