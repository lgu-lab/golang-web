package datamap

import (
	"sync"

	"internal/entities"
	"internal/log"
)

// Structure definition
type StudentDataMap struct {
	dataMap map[int]entities.Student // the map to store Student entities
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
		dataMap: make(map[int]entities.Student),
		lock:    sync.RWMutex{},
	}
}

func (this *StudentDataMap) Read(id int) *entities.Student {
	log.Debug("StudentDataMap - read(%d) ", id)
	this.lock.RLock()
	defer this.lock.RUnlock()
	student, exists := this.dataMap[id]
	if exists {
		return &student
	} else {
		return nil
	}
}
func (this *StudentDataMap) Exists(id int) bool {
	log.Debug("StudentDataMap - exists(%d) ", id)
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, exists := this.dataMap[id]
	return exists
}

func (this *StudentDataMap) Write(student entities.Student) {
	log.Debug("StudentDataMap - write(%+v) ", student)
	this.lock.Lock()
	defer this.lock.Unlock()
	this.dataMap[student.Id] = student
}

func (this *StudentDataMap) Remove(id int) {
	log.Debug("StudentDataMap - remove(%d) ", id)
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.dataMap, id) // delete in map
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
