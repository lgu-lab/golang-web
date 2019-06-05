package memdb

import (
	"sync"

	"internal/entities"
	"internal/log"
)

// Structure definition
type StudentDataMap struct {
	dataMap map[string]entities.Student // the map to store Student entities
	lock    sync.RWMutex
}

var studentDataOnce sync.Once
var studentDataMap  StudentDataMap

func GetStudentDataMap() *StudentDataMap {
	log.Debug("StudentDataMap - GetStudentDataMap() ")
	
	// From Golang doc :
	// "func (o *Once) Do(f func())"
	// "Do" calls the function "f" if and only if Do is being called for the first time for this instance of Once. 
	// In other words, given "var once Once" if "once.Do(f)" is called multiple times,
	// only the first call will invoke f, even if f has a different value in each invocation. 
	// A new instance of Once is required for each function to execute. 
	studentDataOnce.Do(newStudentDataMap) // called only 1 time
	return &studentDataMap
}

func newStudentDataMap() {
	log.Debug("StudentDataMap - newStudentDataMap() ***** ")
	studentDataMap = StudentDataMap{
		dataMap: make(map[string]entities.Student),
		lock:    sync.RWMutex{},
	}
}

func (this *StudentDataMap) Read(id int) *entities.Student {
	key := buildKey(id)
	log.Debug("StudentDataMap - read '%s' ", key)
	this.lock.RLock()
	defer this.lock.RUnlock()
	student, exists := this.dataMap[key]
	if exists {
		return &student
	} else {
		return nil
	}
}
func (this *StudentDataMap) Exists(id int) bool {
	key := buildKey(id)
	log.Debug("StudentDataMap - exists '%s' ", key)
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, exists := this.dataMap[key]
	return exists
}

func (this *StudentDataMap) Write(student entities.Student) {
	key := buildKey(student.Id)
	log.Debug("StudentDataMap - write '%s' : %+v ", key, student)
	this.lock.Lock()
	defer this.lock.Unlock()
	this.dataMap[key] = student
}

func (this *StudentDataMap) Remove(id int) {
	key := buildKey(id)
	log.Debug("StudentDataMap - remove '%s' ", key)
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.dataMap, key) // delete in map
}

func (this *StudentDataMap) Values() []entities.Student {
	this.lock.Lock()
	defer this.lock.Unlock()
	var a = make([]entities.Student, len(this.dataMap))
	i := 0
	for _, v := range this.dataMap {
		a[i] = v
		i++
	}
	return a
}
