package persistence

import (
	"internal/persistence/dao"		
	"internal/persistence/impl/mem"
)

func GetLanguageDAO() dao.LanguageDAO {
	daoImpl := memdb.NewLanguageDAOMemory() // DAO implementation 
	return &daoImpl
}
