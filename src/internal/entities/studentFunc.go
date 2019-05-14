package entities

import 	(
	"sort"
)

func NewStudent() Student {
	// new Student with default values ( 'zero values' )
	return Student{}
}

func SortStudentsById(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Id < students[j].Id
	})
}
