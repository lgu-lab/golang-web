package main

import (
	"fmt"
	"strings"
	"internal/entities"
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
    
    language := entities.NewLanguage()
    language.Code = "AA"
    language.Name = "Bbbbbb"
    log.Debug("language (v)  : %v ", language ) 
    log.Debug("language (+v) : %+v ", language )

    variadicArgs(1, 2, 3)
    variadicArgs("a", 2, true)
    
    fmt.Println(buildKey("a", "bb", "ccc"))
    fmt.Println(buildKey("a", 2, true)) 
}

func f(format string, v ...interface{}) {
    fmt.Printf(format, v...)
}

func variadicArgs(args ...interface{}) {
	fmt.Printf("len(args) = %d \n", len(args) )
	for i, arg := range args {
		fmt.Printf("%d : %v \n", i, arg )
	}
}

func buildKey(args ...interface{}) string {
	var b strings.Builder
	for i, arg := range args {
		if ( i > 0 ) {
			b.WriteString("|")
		}
		// b.WriteString(string(arg))
		// b.WriteString(arg.(string)) // panic: interface conversion: interface {} is int, not string
		fmt.Fprintf(&b, "%v", arg)
		
	}
	return b.String()
}
