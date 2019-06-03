package rest

import (
	"net/http"

	controllers "internal/web/rest/controllers"	
)

func InitRestControllers() {

	apiRoot := "/api/v1" 
	
	// Specific Paths with specific controllers 

	languageRestController := controllers.NewLanguageRestController(apiRoot + "/language") 
	http.HandleFunc( languageRestController.URI(), languageRestController.Process)
	http.HandleFunc( languageRestController.URI()+"/", languageRestController.Process)

}
