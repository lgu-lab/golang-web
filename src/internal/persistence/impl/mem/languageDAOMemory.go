package memdb

import (
	"internal/entities"
	"internal/log"
	"internal/persistence/dao"	
)

//-----------------------------------------------------
// Langage DAO with "in memory data" (based on a map)
//-----------------------------------------------------

// Structure definition
type LanguageDAOMemory struct {
	dataMap * LanguageDataMap
}

// Check interface implementation is valid
var _ dao.LanguageDAO = (*LanguageDAOMemory)(nil)

// Pseudo-construtor
func NewLanguageDAOMemory() LanguageDAOMemory {
	log.Debug("NewLanguageDAOMemory()")
	return LanguageDAOMemory{
		dataMap: GetLanguageDataMap(),
	}
}

func (this *LanguageDAOMemory) FindAll() []entities.Language {
	log.Debug("LanguageDAOMemory - FindAll() ")
	all := this.dataMap.Values()
	entities.SortLanguageByName(all)
	return all 
}

func (this *LanguageDAOMemory) Find(code string) *entities.Language {
	log.Debug("LanguageDAOMemory - Find(%s) ", code)
	return this.dataMap.Read(code)
}

func (this *LanguageDAOMemory) Exists(code string) bool {
	log.Debug("LanguageDAOMemory - Exists(%s) ", code)
	exists := this.dataMap.Exists(code)
	log.Debug("LanguageDAOMemory - Exists(%s, %t) : ", code, exists)
	return exists
}

func (this *LanguageDAOMemory) Create(language entities.Language) bool {
	log.Debug("LanguageDAOMemory - Create(%s) ", language.Code)
	if this.Exists(language.Code) {
		return false // already exists => cannot create
	} else {
		this.dataMap.Write(language)
		return true // not found => created
	}
}

func (this *LanguageDAOMemory) Delete(code string) bool {
	log.Debug("LanguageDAOMemory - Delete(%s) ", code)
	if this.Exists(code) {
		this.dataMap.Remove(code) // delete in map
		return true  // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *LanguageDAOMemory) Update(language entities.Language) bool {
	log.Debug("LanguageDAOMemory - Update(%s) ", language.Code)
	if this.Exists(language.Code) {
		this.dataMap.Write(language) // update in map
		return true  // found and updated
	} else {
		return false // not found => not updated
	}
}
