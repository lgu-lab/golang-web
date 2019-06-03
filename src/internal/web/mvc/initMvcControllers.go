package mvc

import (
	"net/http"
	
	controllers "internal/web/mvc/controllers"  // keep full path for "go build"
)

func InitMvcControllers() {

	// Specific Paths with specific controllers 

	languageController := controllers.NewLanguageController() 
	http.HandleFunc("/language/list", languageController.ListHandler)
	http.HandleFunc("/language/form", languageController.FormHandler)

	studentController := controllers.NewStudentController() 
	http.HandleFunc("/student/list", studentController.ListHandler )
	http.HandleFunc("/student/form", studentController.FormHandler )

}