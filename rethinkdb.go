package main

import (
	"sort"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/afeldman/go-util/env"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	Session *r.Session
	once    sync.Once
)

func getorcreate() r.Term {
	if Session == nil {
		url_session()
	}

	databasename := env.GetEnvOrDefault("UNIQUE_URL_DB", "hash_url")
	tablename := env.GetEnvOrDefault("UNIQUE_URL_DB_TABLES", "urls")

	var dbnames []string

	dbcur, err := r.DBList().Run(Session)
	checkError(err)
	err = dbcur.All(&dbnames)
	checkError(err)
	defer dbcur.Close()

	if !contains(dbnames, databasename) {
		log.Debugln("createDB: ", databasename)
		r.DBCreate(databasename).Run(Session)
	}

	var tbnames []string
	tbcur, err := r.DB(databasename).TableList().Run(Session)
	checkError(err)
	err = tbcur.All(&tbnames)
	checkError(err)
	defer tbcur.Close()

	if !contains(tbnames, tablename) {
		r.DB(databasename).TableCreate(tablename).Run(Session)
	}

	return r.DB(databasename).Table(tablename)
}

func setdata(entry *DB_entry) {
	term := getorcreate()

	log.Debugln("try to store ", entry.Hash)

	term.Insert(&entry).Run(Session)
}

func getdata(hash string) *DB_entry {
	term := getorcreate()

	log.Debugln("get data of key ", hash)

	var entry *DB_entry

	cur, err := term.Get(hash).Run(Session)
	checkError(err)
	if cur.IsNil() {
		return nil
	}
	defer cur.Close()

	err = cur.One(&entry)
	checkError(err)

	return entry
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func url_session() *r.Session {

	once.Do(func() {
		db_address := env.GetEnvOrDefault("UNIQUE_URL_DB_ADDRESS", "localhost:28015")

		session, err := r.Connect(r.ConnectOpts{
			Address: db_address, // endpoint without http
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Infoln("Database on ", db_address)

		Session = session
	})

	return Session
}
