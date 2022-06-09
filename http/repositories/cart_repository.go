package repositories

import (
	mongoEntity "example/databases/entities/mongo"
	"example/http/interfaces"

	mgo "gopkg.in/mgo.v2"
)

type cartRepository struct {
	dbSession *mgo.Session
	database  string
}

const (
	collectionCart = "cart"
)

func NewCartRepository(sess *mgo.Session, database string) interfaces.CartInterface {
	return &cartRepository{sess, database}
}

func (model *cartRepository) GetAllCart() ([]*mongoEntity.Cart, error) {

	ds := model.dbSession.Copy()
	defer ds.Close()

	table := ds.DB(model.database).C(collectionCart)
	listCart := []*mongoEntity.Cart{}
	err := table.Find(nil).All(&listCart)

	if err != nil {
		return listCart, err
	}
	return listCart, nil
}
