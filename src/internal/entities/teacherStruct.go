package entities

import "fmt"

type Teacher struct {
	Id        int 
	FirstName string
	LastName  string
}

func NewTeacher() Teacher {
	// new Teacher with default values ( 'zero values' )
	return Teacher{}
}

func (this Teacher) String() string {
    return fmt.Sprintf(
    	"[%d : %s, %s]", 
    	this.Id, 
    	this.FirstName, 
    	this.LastName)
}