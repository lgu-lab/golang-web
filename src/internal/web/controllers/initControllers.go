package controllers

import (
	"net/http"
	
	//"internal/web/controllers"  // keep full path for "go build"
)

func InitControllers() {

	// Specific Paths with specific controllers 

	languageController := NewLanguageController() 
	http.HandleFunc("/language/list", languageController.ListHandler)
	http.HandleFunc("/language/form", languageController.FormHandler)

	studentController := NewStudentController() 
	http.HandleFunc("/student/list", studentController.ListHandler )
	http.HandleFunc("/student/form", studentController.FormHandler )

}