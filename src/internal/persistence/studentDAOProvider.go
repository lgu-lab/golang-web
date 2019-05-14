package persistence

import (
	"internal/persistence/dao"		
	"internal/persistence/impl/mem"
)

func GetStudentDAO() dao.StudentDAO {
	daoImpl := memdb.NewStudentDAOMemory() // DAO implementation 
	return &daoImpl
}
