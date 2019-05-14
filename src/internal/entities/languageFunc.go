package entities

import 	(
	"sort"
)

func NewLanguage() Language {
	// new structure with default values ( 'zero values' )
	return Language{}
}

func NewLanguageInit(code string, name string) Language {
	// new structure with specific values 
	return Language{
		Code: code,
		Name: name,
	}
}

func SortLanguageByName(entities []Language) {
	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Name < entities[j].Name
	})
}

func SortLanguageByCode(entities []Language) {
	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Code < entities[j].Code
	})
}
