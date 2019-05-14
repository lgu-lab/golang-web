package dao

import (
	"database/sql"
	"internal/log"
	"internal/entities"
	"internal/persistence/dao"
)

//-----------------------------------------------------
// Langage DAO : SQL implem
//-----------------------------------------------------

// TaskDAOSQL is the sql implementation of the TaskDAO
type LanguageDAOSQL struct {
	sqlSession *sql.DB
}

// Check interface implementation is valid
var _ dao.LanguageDAO = (*LanguageDAOSQL)(nil)

// Pseudo-construtor
func NewLanguageDAOSQL(sqlSession *sql.DB) LanguageDAOSQL {
	log.Debug("NewLanguageDAOSQL()")
	return LanguageDAOSQL {
		sqlSession: sqlSession,
	}
}

// Interface functions implementation 
// TODO
func (this *LanguageDAOSQL) FindAll() []entities.Language {
	return nil
}
func (this *LanguageDAOSQL) Find(code string) *entities.Language {
	return nil
}
func (this *LanguageDAOSQL) Exists(code string) bool {
	return false
}
func (this *LanguageDAOSQL) Create(language entities.Language) bool {
	return false
}
func (this *LanguageDAOSQL) Delete(code string) bool {
	return false
}
func (this *LanguageDAOSQL) Update(language entities.Language) bool {
	return false
}
