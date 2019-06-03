package rest

import (
	"net/http"
	"strings"

	"internal/entities"
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
	case "POST":
	    this.processPOST(w,r)
	case "PUT":
	    this.processPUT(w,r)
	case "DELETE":
		this.processDELETE(w, r)
		
	default:
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
		ReplyNotFound(w)
	}
}

func (this *LanguageRestController) processPOST(w http.ResponseWriter, r *http.Request) {
	log.Debug("processPOST - URL path : " + r.URL.Path)

	// 
	entity := entities.NewLanguage() 
	//err := json.Unmarshal(jsonData, &language)
	err := ReadJSON(&entity, r)
	if err != nil {
		ReplyBadRequest(w)
	}

	if ( this.dao.Create(entity) ) {
		ReplyCreated(w)
	} else {
		ReplyNotCreated(w)
	}
}

func (this *LanguageRestController) processPUT(w http.ResponseWriter, r *http.Request) {
	log.Debug("processPUT - URL path : " + r.URL.Path)
	
	entity := entities.NewLanguage() 
	err := ReadJSON(&entity, r)
	if err != nil {
		ReplyBadRequest(w)
	}

	if ( this.dao.Update(entity) ) {
		ReplyUpdated(w)
	} else {
		ReplyNotUpdated(w)
	}
}

func (this *LanguageRestController) processDELETE(w http.ResponseWriter, r *http.Request) {
	log.Debug("processDELETE - URL path : " + r.URL.Path)
	
	// get key
	parts := strings.Split(r.URL.Path, "/")
	n := this.uriPartsCount
	k1 := parts[n] 
//	k2 := parts[n+1]

	// get data
	log.Debug("deleteById - dao.Delete : " + k1 )
	if ( this.dao.Delete(k1) ) {
		ReplyDeleted(w)
	} else {
		ReplyNotDeleted(w)
	}
}
