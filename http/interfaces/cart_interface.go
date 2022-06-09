package interfaces

import mongoEntity "example/databases/entities/mongo"

type CartInterface interface {
	GetAllCart() ([]*mongoEntity.Cart, error)
}
