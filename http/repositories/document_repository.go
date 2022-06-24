package repositories

import (
	"apl-ezrx-api/src/Services/Cart.api/http/interfaces"

	mgo "gopkg.in/mgo.v2"
)

type documentRepository struct {
	dbSession *mgo.Session
	database  string
}

const (
	collectionDocument = "document"
)

func NewDocumentRepository(sess *mgo.Session, database string) interfaces.DocumentInterface {
	return &documentRepository{sess, database}
}

// begin this is jus an example we SHOLUD NOT pu a struct here
type DummyStructMongoDBEntity struct {
	Nama   string
	Alamat string
}

// end this is jus an example we SHOLUD NOT pu a struct here

func (repo *documentRepository) CreateBulk(documents []*DummyStructMongoDBEntity) error {

	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()

	interfaceTobeSave := make([]interface{}, len(documents))
	table := ds.DB(repo.database).C(collectionDocument)

	for i := 0; i < len(documents); i++ {
		interfaceTobeSave = append(interfaceTobeSave, documents[i])
	}

	if err = table.Insert(interfaceTobeSave...); err != nil {
		return err
	}

	return err
}
