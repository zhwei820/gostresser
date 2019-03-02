package config

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/json-iterator/go"
	"github.com/narup/gmgo"
	"log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Dbconfig = gmgo.DbConfig{
	HostURL: "mongodb://localhost:27017/userdb",
	DBName:  "userdb"}

var TestDB gmgo.Db

func init() {
	setupDB()
}
func setupDB() {
	if err := gmgo.Setup(gmgo.DbConfig{
		HostURL: "mongodb://localhost:27017/userdb",
		DBName:  "userdb"}); err != nil {
		log.Fatalf("Database connection error : %s.\n", err)
		return
	}

	newDb, err := gmgo.Get("userdb")
	if err != nil {
		log.Fatalf("Db connection error : %s.\n", err)
	}
	TestDB = newDb
}

// create index
func EnsureIndex(colname string, indexes []mgo.Index) {
	session := TestDB.Session()
	defer session.Close()

	c := session.Session.DB(Dbconfig.DBName).C(colname)
	for _, index := range indexes {
		err := c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}
	}

	idxes, _ := c.Indexes()
	for idx, item := range idxes {
		println(fmt.Sprintf("%v, %v", idx, item))
	}
}
