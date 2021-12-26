package main

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

func post_unique(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	// declaring new post of type Post
	var request Request

	// reads the JSON value and decodes it into a Go value
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	u, err := url.ParseRequestURI(request.Url)
	checkError(err)

	log.Debugln("store: ", request.Url)
	h := sha3.Sum256([]byte(request.Url))

	var entry DB_entry

	entry.Hash = hex.EncodeToString(h[:])
	entry.Url = u
	entry.ConnectCount = 0

	log.Debugln("store data: ", entry.Hash)
	setdata(&entry)

	result, err := json.Marshal(entry)
	checkError(err)

	res.WriteHeader(http.StatusOK)

	res.Write(result)
}

type answer struct {
	Url string `json:"url"`
}

func get_unique(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	hash := vars["hash"]

	entry := getdata(hash)

	var ans answer
	if entry == nil {
		ans.Url = "none"
		res.WriteHeader(http.StatusBadRequest)
	} else {
		ans.Url = entry.Url.String()
		res.WriteHeader(http.StatusOK)
	}

	result, err := json.Marshal(ans)
	checkError(err)

	log.Debugln(string(result))

	res.Write(result)
}
