package datamap

import (
	"fmt"
	"testing"
	
	"internal/entities"
)

func Test1(t *testing.T) {
	fmt.Println(t.Name()) // prints function name

	dataMap1 := GetLanguageDataMap()
	dataMap2 := GetLanguageDataMap()
	dataMap3 := GetLanguageDataMap()
	
	if dataMap1 != dataMap2 {
		t.Error("Not a single map !")
	}
	if dataMap1 != dataMap3 {
		t.Error("Not a single map !")
	}

	// check same instance :
	// write in map 1
	language := entities.NewLanguage()
	language.Code = "JA"
	language.Name = "Java"
	dataMap1.write(language)

	// read in map 2
	if ! dataMap2.exists("JA") {
		t.Error("Language not found in map 2")
	}
}

func Test2(t *testing.T) {
	fmt.Println(t.Name()) // prints function name

	dataMap := GetLanguageDataMap()
	//exists := dataMap.exists("JA")
	if dataMap.exists("JA") {
		fmt.Println("Language found => remove")
		dataMap.remove("JA")
	}
	
	if dataMap.exists("JA") {
		t.Error("Language found! (expected 'not found')")
	}
		
	language := entities.NewLanguage()
	language.Code = "JA"
	language.Name = "Java"
	dataMap.write(language)
	if ! dataMap.exists("JA") {
		t.Error("Language not found! (expected 'found')")
	}
	//			t.Errorf("Error just for tests")
	//			t.Fail()
	//t.Error("Fake error, just for test")
}
