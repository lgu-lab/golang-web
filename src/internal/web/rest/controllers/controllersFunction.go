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

// GET 
func ReplyFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
func ReplyNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// DELETE
func ReplyDeleted(w http.ResponseWriter) {
	// StatusNoContent = 204 // RFC 7231, 6.3.5
	// The server successfully processed the request, but is not returning any content
	w.WriteHeader(http.StatusNoContent)
}
func ReplyNotDeleted(w http.ResponseWriter) {
	// StatusNotFound = 404 // RFC 7231, 6.5.4
	// Not deleted so supposed "not found"
	w.WriteHeader(http.StatusNotFound)
}

func ReplyNotImplemented(w http.ResponseWriter) {
	// StatusNotImplemented = 501 // RFC 7231, 6.6.2
	w.WriteHeader(http.StatusNotImplemented)
}
