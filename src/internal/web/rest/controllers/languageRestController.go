package restcontrollers

import (
	"encoding/json"
	"net/http"

//	"internal/entities"
	"internal/log"
	"internal/persistence"
	"internal/persistence/dao"
	"internal/webutil"
)

type LanguageRestController struct {
	uri string
	dao dao.LanguageDAO // DAO interface (abstract)
}

func NewLanguageRestController(uri string) LanguageRestController {
	return LanguageRestController {
		uri: uri,
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
		webutil.ErrorPage(w, "Method "+r.Method+" is not supported")
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (this *LanguageRestController) getById(w http.ResponseWriter, r *http.Request) {
	log.Debug("getById - URL path : " + r.URL.Path)
	// get data
	// data := this.dao.Find(code string)
	// return data
	w.Header().Set("Content-Type", "application/json")
	//w.Write("{\"Debug\":\"getById\"")
}
