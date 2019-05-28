package main

import (
	"fmt"
	"encoding/json"	
	"internal/entities"
)

func main() {
    fmt.Println("JSON Marshal/Unmarshal")

	language := entities.NewLanguage() 
	language.Code = "JA"
	language.Name = "Java"
    
//    fmt.Printf("language : %s  \n", language.String() )
//    fmt.Printf("language : %+v  \n", language )
    fmt.Printf("language : %v  \n", language )
    
    fmt.Println("-----")
    
    fmt.Printf("json.Marshal(language)...\n" )
    // func Marshal(v interface{}) ([]byte, error)
    jsonData, err := json.Marshal(language)
	if err != nil {
	    fmt.Printf("Marshal error : " + language.String() )
		return
	}
    fmt.Printf("JSON : %x \n", jsonData )
    fmt.Printf("JSON : %s \n", string(jsonData) )
    
    fmt.Println("-----")
    
    fmt.Printf("json.Unmarshal(jsonData, &language2)...\n" )
	language2 := entities.NewLanguage() 
	err2 := json.Unmarshal(jsonData, &language2)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
    fmt.Printf("language2 : %v  \n", language2 )
    
}

