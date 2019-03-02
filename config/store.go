package config

import (
	"github.com/globalsign/mgo/bson"
	"github.com/narup/gmgo"
	"github.com/pkg/errors"
)

//
func BaseConfManager() *ManageBaseConf {
	return new(ManageBaseConf)
}

type ManageBaseConf struct {
}

// create
func (manager *ManageBaseConf) Create(baseConf BaseConf) (bson.ObjectId, error) {
	session := TestDB.Session()
	defer session.Close()

	baseConf.Id = bson.NewObjectId()
	err := session.Save(baseConf)
	return baseConf.Id, err
}

// Save
func (manager *ManageBaseConf) Save(baseConf BaseConf) (bson.ObjectId, error) {
	session := TestDB.Session()
	defer session.Close()

	err := session.Update(gmgo.Q{"_id": baseConf.Id}, baseConf)
	return baseConf.Id, err
}

// Update
func (manager *ManageBaseConf) Update(q gmgo.Q, update interface{}) error {
	session := TestDB.Session()
	defer session.Close()

	err := session.Session.DB(Dbconfig.DBName).C(BaseConf{}.CollectionName()).Update(q, update)
	return err
}

// read
func (manager *ManageBaseConf) FindOne(baseConfId string) (*BaseConf, error) {
	session := TestDB.Session()
	defer session.Close()

	baseConf := new(BaseConf)
	err := session.FindByID(baseConfId, baseConf)
	return baseConf, err
}

// RemoveOne
func (manager *ManageBaseConf) RemoveOne(baseConfId string) (*BaseConf, error) {
	session := TestDB.Session()
	defer session.Close()

	baseConf := new(BaseConf)
	err := session.Remove(gmgo.Q{"_id": bson.ObjectIdHex(baseConfId)}, baseConf)
	return nil, err
}

// read all
func (manager *ManageBaseConf) FindAll(q gmgo.Q) ([]BaseConf, error) {
	session := TestDB.Session()
	defer session.Close()
	var searchResults []BaseConf
	col := session.Session.DB(Dbconfig.DBName).C(BaseConf{}.CollectionName())

	err := col.Find(q).Skip(0).All(&searchResults)
	return searchResults, err
}

// filter One
func (manager *ManageBaseConf) FilterOne(q gmgo.Q) (*BaseConf, error) {
	session := TestDB.Session()
	defer session.Close()

	var searchResults []BaseConf
	col := session.Session.DB(Dbconfig.DBName).C(BaseConf{}.CollectionName())
	var err error
	err = col.Find(q).Skip(0).Limit(1).Sort().All(&searchResults)
	if len(searchResults) == 0 {
		return nil, errors.Errorf("not found")
	}
	return &searchResults[0], err
}

// FindPaged
func (manager *ManageBaseConf) FindPaged(q gmgo.Q, skip int, limit int, sort []string) (int, []BaseConf, error) {
	session := TestDB.Session()
	defer session.Close()

	var searchResults []BaseConf
	col := session.Session.DB(Dbconfig.DBName).C(BaseConf{}.CollectionName())
	var err error
	err = col.Find(q).Skip(skip).Limit(limit).Sort(sort...).All(&searchResults)
	if limit <= 0 {
		err = col.Find(q).Skip(skip).All(&searchResults)
	}
	count, err := col.Find(q).Count()

	return count, searchResults, err
}

// Count
func (manager *ManageBaseConf) Count(q gmgo.Q) (int, error) {
	session := TestDB.Session()
	defer session.Close()

	col := session.Session.DB(Dbconfig.DBName).C(BaseConf{}.CollectionName())
	count, err := col.Find(q).Count()

	return count, err
}
