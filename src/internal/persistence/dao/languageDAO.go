package dao

import (
	"internal/entities"
)

//-----------------------------------------------------
// Langage DAO INTERFACE
//-----------------------------------------------------
type LanguageDAO interface {
	FindAll() []entities.Language
	Find(code string) *entities.Language
	Exists(code string) bool
	Create(language entities.Language) bool
	Delete(code string) bool
	Update(language entities.Language) bool 
}
