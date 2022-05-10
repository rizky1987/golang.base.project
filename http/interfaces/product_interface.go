package interfaces

import (
	entity "example/databases/entities/sql"

	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(transaction *gorm.DB, data entity.Product) (*entity.Product, error)
	GetProductByCode(code string) (*entity.Product, error)
}
