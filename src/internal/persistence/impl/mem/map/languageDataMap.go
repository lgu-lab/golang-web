package datamap

import (
	"log"
	"sync"

	"internal/entities"
)

// Structure definition
type LanguageDataMap struct {
	dataMap map[string]entities.Language // the map to store Language entities
	lock    sync.RWMutex
}

var languageDataOnce sync.Once
var languageDataMap  LanguageDataMap

func GetLanguageDataMap() *LanguageDataMap {
	log.Printf("LanguageDataMap - GetLanguageDataMap() ")
	
	// From Golang doc :
	// "func (o *Once) Do(f func())"
	// "Do" calls the function "f" if and only if Do is being called for the first time for this instance of Once. 
	// In other words, given "var once Once" if "once.Do(f)" is called multiple times,
	// only the first call will invoke f, even if f has a different value in each invocation. 
	// A new instance of Once is required for each function to execute. 
	languageDataOnce.Do(newLanguageDataMap) // called only 1 time
	return &languageDataMap
}

func newLanguageDataMap() {
	log.Printf("LanguageDataMap - newLanguageDataMap() ***** ")
	languageDataMap = LanguageDataMap{
		dataMap: make(map[string]entities.Language),
		lock:    sync.RWMutex{},
	}
}

func (this *LanguageDataMap) Read(code string) *entities.Language {
	log.Printf("LanguageDataMap - read(%s) ", code)
	this.lock.RLock()
	defer this.lock.RUnlock()
	language, exists := this.dataMap[code]
	if exists {
		return &language
	} else {
		return nil
	}
}
func (this *LanguageDataMap) Exists(code string) bool {
	log.Printf("LanguageDataMap - exists(%s) ", code)
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, exists := this.dataMap[code]
	return exists
}

func (this *LanguageDataMap) Write(language entities.Language) {
	log.Printf("LanguageDataMap - write(%s) ", language.String())
	this.lock.Lock()
	defer this.lock.Unlock()
	this.dataMap[language.Code] = language
}

func (this *LanguageDataMap) Remove(code string) {
	log.Printf("LanguageDataMap - remove(%s) ", code)
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.dataMap, code) // delete in map
}

func (this *LanguageDataMap) Values() []entities.Language {
	this.lock.Lock()
	defer this.lock.Unlock()
	var a = make([]entities.Language, len(this.dataMap))
	i := 0
	for _, v := range this.dataMap {
		a[i] = v
		i++
	}
	return a
}
