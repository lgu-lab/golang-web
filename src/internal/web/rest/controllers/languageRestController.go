package restcontrollers

import (
//	"encoding/json"
	"net/http"
	"strings"

//	"internal/entities"
	"internal/log"
	"internal/persistence"
	"internal/persistence/dao"
)

type LanguageRestController struct {
	uri string
	uriPartsCount int
	dao dao.LanguageDAO // DAO interface (abstract)
}

func NewLanguageRestController(uri string) LanguageRestController {
	parts := strings.Split(uri, "/")
	return LanguageRestController {
		uri: uri,
		uriPartsCount: len(parts),
		dao: persistence.GetLanguageDAO(), // DAO implementation
	}
}

func (this *LanguageRestController) URI() string {
	return this.uri
}

func (this *LanguageRestController) Process(w http.ResponseWriter, r *http.Request) {
	log.Debug("Process - URL path : " + r.URL.Path)

	switch r.Method {
	case "GET":
		this.processGET(w, r)
		//	case "POST":
		//	    this.processPost(w,r)
	default:
//		webutil.ErrorPage(w, "Method "+r.Method+" is not supported")
		http.Error(w, "", http.StatusBadRequest)
	}
}

func (this *LanguageRestController) processGET(w http.ResponseWriter, r *http.Request) {
	log.Debug("processGET - URL path : " + r.URL.Path)
	if ( len(r.URL.Path) > len(this.uri) ) {
		// more than the URI itself => id	
		this.getById(w,r)	
	} else {
		// just the URI itself => no id
		this.getAll(w, r)
	}
}

func (this *LanguageRestController) getAll(w http.ResponseWriter, r *http.Request) {
	log.Debug("getAll - URL path : " + r.URL.Path)
	// get data
	data := this.dao.FindAll()

//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonData)
	
	WriteJSON(w, data) 
}

func (this *LanguageRestController) getById(w http.ResponseWriter, r *http.Request) {
	log.Debug("getById - URL path : " + r.URL.Path)
	
	// get key
	parts := strings.Split(r.URL.Path, "/")
	n := this.uriPartsCount

//	log.Debug("getById - parts : %v (initial count = %d)", parts, this.uriPartsCount )

	k1 := parts[n] 
//	k2 := parts[n+1]

	// get data
	log.Debug("getById - dao.Find : " + k1 )
	data := this.dao.Find(k1)
	if ( data != nil ) {
		WriteJSON(w, data)
	} else {
		ReplyStatusNotFound(w)
	}
}
