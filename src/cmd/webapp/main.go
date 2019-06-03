package main

import (
	"log"
	"net/http"
	
	"internal/utils" 
	"internal/web/mvc"
	"internal/web/rest"
	
//	"webapp/persistence/dao"
//	"webapp/persistence/bolt"
)

const webDir  = "www"
const webPort = ":80"

func serveStaticFile(w http.ResponseWriter, r *http.Request) {		
	s := webDir + r.URL.Path
	log.Print("Static file requested. URL path '" + r.URL.Path + "' --> server file '" + s + "'" )
	http.ServeFile(w, r, s)
}

//func initBoltDB() {
//	log.Print("Init Bolt DB ... ")
//	bolt.BoltStart("boltdb.data")
//	dao.SetDAOImplementation(dao.BOLTDB)	
//}

func main() {
	
	// initBoltDB()
	
	log.Print("Starting server ... ")

	utils.PrintEnv() 

	log.Print("Setting static file handler... ")
	// Set handler to serve static files
	http.HandleFunc("/", serveStaticFile)			

	log.Print("Setting application handlers... ")
	mvc.InitMvcControllers()
	rest.InitRestControllers()

	// Launch the http server
	log.Print("Launching http server (port=" + webPort + ") ... ")
	log.Fatal(http.ListenAndServe(webPort, nil))

}