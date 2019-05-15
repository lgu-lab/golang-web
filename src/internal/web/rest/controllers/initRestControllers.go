package restcontrollers

import (
	"net/http"
	
	//"internal/web/controllers"  // keep full path for "go build"
)

func InitRESTControllers() {

	apiRoot := "/api/v1" 
	
	// Specific Paths with specific controllers 

	languageRestController := NewLanguageRestController(apiRoot + "/language") 
	http.HandleFunc( languageRestController.URI(), languageRestController.Process)
	http.HandleFunc( languageRestController.URI()+"/", languageRestController.Process)


}