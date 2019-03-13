package config

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/json-iterator/go"
	"github.com/micro/go-config"
	"github.com/narup/gmgo"
	"log"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Dbconfig = gmgo.DbConfig{
	HostURL: "mongodb://localhost:27017/testdb",
	DBName:  "testdb",
}

var TestDB gmgo.Db

func init() {
	build_env := os.Getenv("BUILD_ENV")
	if build_env == "not_dev" {
		config.LoadFile("./config/config.yaml")
		config.Scan(&Dbconfig)
	}
	setupDB()
}
func setupDB() {
	if err := gmgo.Setup(Dbconfig); err != nil {
		log.Fatalf("Database connection error : %s.\n", err)
		return
	}

	newDb, err := gmgo.Get("testdb")
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
