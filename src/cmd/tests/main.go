package main

import (
	"fmt"
	"internal/log"
)

func main() {
    fmt.Println("hello world")

    fmt.Println("hello world")
    
    fmt.Printf("Boolean : %t %t \n", true, false)

    f("Boolean : %t %t \n", true, false)
    
    s := "AZERTY"
    log.Info("aaaaa" )
    log.Info("bbb" + s )
    log.Info("bb '%s' cc '%d' dd '%f'", "xxx", 123, 89.777 )
    log.Info("---" )

    log.Debug("aaaaa" )
    log.Debug("bb '%s' cc '%d' dd '%f'", "xxx", 123, 89.777 )
}

func f(format string, v ...interface{}) {
    fmt.Printf(format, v...)
}
