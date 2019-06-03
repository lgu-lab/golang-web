package rest

import (
	"net/http"
	"encoding/json"

)

func WriteJSON(w http.ResponseWriter, data interface{}) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ReadJSON(data interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(data)
}

// GET 
func ReplyFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
func ReplyNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// CREATE
func ReplyCreated(w http.ResponseWriter) {
	// StatusCreated = 201 // RFC 7231, 6.3.2
	// The 201 (Created) status code indicates that the request has been
	// fulfilled and has resulted in one or more new resources being created. 
	w.WriteHeader(http.StatusCreated)
}
func ReplyNotCreated(w http.ResponseWriter) {
	// StatusConflict = 409 // RFC 7231, 6.5.8
	// The 409 (Conflict) status code indicates that the request could not
	// be completed due to a conflict with the current state of the target resource.
	w.WriteHeader(http.StatusConflict)
}

// UPDATE 
func ReplyUpdated(w http.ResponseWriter) {
	// StatusOK = 200 // RFC 7231, 6.3.1
	w.WriteHeader(http.StatusOK)
}
func ReplyNotUpdated(w http.ResponseWriter) {
	// StatusNotFound = 404 // RFC 7231, 6.5.4
	// Not updated so supposed "not found"
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

// ERRORS
func ReplyBadRequest(w http.ResponseWriter) {
	// StatusBadRequest = 400 // RFC 7231, 6.5.1
	w.WriteHeader(http.StatusBadRequest)
}

func ReplyNotImplemented(w http.ResponseWriter) {
	// StatusNotImplemented = 501 // RFC 7231, 6.6.2
	w.WriteHeader(http.StatusNotImplemented)
}
