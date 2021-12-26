package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/afeldman/go-util/env"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

type DB_entry struct {
	Hash         string `rethinkdb:"id"`
	ConnectCount uint8  `rethinkdb:"counter"`
	Url          *url.URL
}

// defining Post entity
type Request struct {
	Url string `json:"url"`
}

func main() {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file")
	}

	// set database connection
	url_session()

	r := mux.NewRouter()

	r.HandleFunc("/{hash}", get_unique).Methods("GET")
	r.HandleFunc("/", post_unique).Methods("POST")
	r.HandleFunc("/{hash}", delete_unique).Methods("DELETE")

	server_Address := env.GetEnvOrDefault("UNIQUE_URL_ADDRESS", "localhost")
	server_Port := env.GetEnvOrDefault("UNIQUE_URL_PORT", "1103")

	srv := &http.Server{
		Handler: r,
		Addr:    server_Address + ":" + server_Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infoln("Server starts on " + server_Address + ":" + server_Port)

	log.Fatalln(srv.ListenAndServe())
}
