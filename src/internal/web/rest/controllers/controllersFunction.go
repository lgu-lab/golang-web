package restcontrollers

import (
	"net/http"
	"encoding/json"
	
	//"internal/web/controllers"  // keep full path for "go build"
)

func InitRESTControllers() {

	apiRoot := "/api/v1" 
	
	// Specific Paths with specific controllers 

	languageRestController := NewLanguageRestController(apiRoot + "/language") 
	http.HandleFunc( languageRestController.URI(), languageRestController.Process)
	http.HandleFunc( languageRestController.URI()+"/", languageRestController.Process)

}

func WriteJSON(w http.ResponseWriter, data interface{}) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ReplyStatusNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

