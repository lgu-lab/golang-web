package main

import (
	"fmt"
	"internal/entities"
	"internal/log"
	"strings"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("hello world")

	fmt.Printf("Boolean : %t %t \n", true, false)

	f("Boolean : %t %t \n", true, false)

	s := "AZERTY"
	log.Info("aaaaa")
	log.Info("bbb" + s)
	log.Info("bb '%s' cc '%d' dd '%f'", "xxx", 123, 89.777)
	log.Info("---")

	log.Debug("aaaaa")
	log.Debug("bb '%s' cc '%d' dd '%f'", "xxx", 123, 89.777)
	log.Debug("All with 'percent v' : %v | %v | %v ", "xxx", 123, 89.777)

	language := entities.NewLanguage()
	language.Code = "AA"
	language.Name = "Bbbbbb"
	log.Debug("language (v)  : %v ", language)
	log.Debug("language (+v) : %+v ", language)

	variadicArgs(1, 2, 3)
	variadicArgs("a", 2, true)

	fmt.Println(buildKey("a", "bb", "ccc"))
	fmt.Println(buildKey("a", 2, true))

	printArgs(getArgs("/a/b/c", "/a/b/c/1"))
	printArgs(getArgs("/a/b/c", "/a/b/c/1/2"))
	printArgs(getArgs("/a/b/c", "/a/b/c"))
	printArgs(getArgs("/a/b/c", "/a/b/c/"))
	printArgs(getArgs("/a/b/c", "/a/b/c//"))
	printArgs(getArgs("/a/b/c", "/a/b/c//x/y"))

	printArgs(getArgs("/a/b/c/", "/a/b/c/1"))
	printArgs(getArgs("/a/b/c/", "/a/b/c/1/2"))
	printArgs(getArgs("/a/b/c/", "/a/b/c"))
	printArgs(getArgs("/a/b/c/", "/a/b/c/"))
	printArgs(getArgs("/a/b/c/", "/a/b/c//"))
	printArgs(getArgs("/a/b/c/", "/a/b/c//x/y"))

	printArgs(getArgs("/",       "/a/b/c/1"))
	printArgs(getArgs("/a/b/c/", "/a/b"))
	printArgs(getArgs("",        "/a/b/c/1"))
	printArgs(getArgs("",        ""))
}
func printArgs(args []string) {
	fmt.Printf("args : %v (size = %d)\n", args, len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf(" . arg[%d] : '%s'\n", i, args[i])
	}
}
func getBaseLength(base string) int {
	baseLength := len(base)
	if baseLength > 0 {
		if base[len(base)-1] == '/' {
			baseLength--
		}
	}
	return baseLength
}
func getArgs(base string, uri string) []string {
	fmt.Println("-----")
	fmt.Println("getArgs('" + base + "','" + uri + "')")

	baseLength := getBaseLength(base)
	if baseLength >= len(uri) || baseLength == 0 {
		return []string{} // Void (no args)
	}
	
	// Keep only the right part (after the base URI)
	keyPart := uri[baseLength:]
	fmt.Println("getArgs : keyPart = '" + keyPart + "'")
	
	// Remove the '/' at the beginning if any
	if keyPart[0] == '/' {
		keyPart = keyPart[1:]
	}
	
	// Split if not void
	if len(keyPart) == 0 {
		return []string{} // Void (no args)
	} else {
		return strings.Split(keyPart, "/")
	}
}

func f(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func variadicArgs(args ...interface{}) {
	fmt.Printf("len(args) = %d \n", len(args))
	for i, arg := range args {
		fmt.Printf("%d : %v \n", i, arg)
	}
}

func buildKey(args ...interface{}) string {
	var b strings.Builder
	for i, arg := range args {
		if i > 0 {
			b.WriteString("|")
		}
		// b.WriteString(string(arg))
		// b.WriteString(arg.(string)) // panic: interface conversion: interface {} is int, not string
		fmt.Fprintf(&b, "%v", arg)

	}
	return b.String()
}
