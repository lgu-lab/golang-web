package memdb

import (
	"internal/entities"
	"internal/log"
	"internal/persistence/dao"	
	"internal/persistence/impl/mem/map"
)

// Structure definition
type StudentDAOMemory struct {
	dataMap * datamap.StudentDataMap
}

// Check interface implementation is valid
var _ dao.StudentDAO = (*StudentDAOMemory)(nil)

// Structure pseudo-construtor
func NewStudentDAOMemory() StudentDAOMemory {
	log.Debug("NewStudentDAOMemory()")
	//dao := StudentDAOMemory{} // structure creation 
	//dao.init() // structure init
	//return dao
	return StudentDAOMemory{
		dataMap: datamap.GetStudentDataMap(),		
	} 
}

func (this *StudentDAOMemory) FindAll() []entities.Student {
	log.Debug("DAO - FindAll() ")
	all := this.dataMap.Values()
	entities.SortStudentsById(all)
	return all
}

func (this *StudentDAOMemory) Find(id int) *entities.Student {
	log.Debug("DAO - Find(%d) ", id)
	return this.dataMap.Read(id)
}

func (this *StudentDAOMemory) Exists(id int) bool {
	log.Debug("DAO - Exists(%d) ", id)
	exists := this.dataMap.Exists(id)
	log.Debug("LanguageDAOMemory - Exists(%s, %t) : ", id, exists)
	return exists
	
}

func (this *StudentDAOMemory) Create(student entities.Student) bool {
	log.Debug("DAO - Create(%d) ", student.Id)
	if this.Exists(student.Id) {
		return false // already exists => cannot create
	} else {
		this.dataMap.Write(student)
		return true // not found => created
	}
	
}

func (this *StudentDAOMemory) Delete(id int) bool {
	log.Debug("DAO - Delete(%d) ", id)
	if this.Exists(id) {
		this.dataMap.Remove(id) // delete in map
		return true  // found and deleted
	} else {
		return false // not found => not deleted
	}
	
}

func (this *StudentDAOMemory) Update(student entities.Student) bool {
	log.Debug("DAO - Update(%d) ", student.Id)
	if this.Exists(student.Id) {
		this.dataMap.Write(student) // update in map
		return true  // found and updated
	} else {
		return false // not found => not updated
	}

}
