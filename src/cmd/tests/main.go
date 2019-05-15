package main

import "fmt"

func main() {
    fmt.Println("hello world")

    fmt.Println("hello world")
    
    fmt.Printf("Boolean : %t %t \n", true, false)

    f("Boolean : %t %t \n", true, false)
    
}

func f(format string, v ...interface{}) {
    fmt.Printf(format, v...)
}
