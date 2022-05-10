package repositories

import (
	"errors"
	entity "example/databases/entities/sql"
	interfaces "example/http/interfaces"

	"gorm.io/gorm"
)

type respositoryUser struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductInterface {
	return &respositoryUser{db}
}

func (ctx respositoryUser) Create(transaction *gorm.DB, data entity.Product) (*entity.Product, error) {
	var err error

	if transaction != nil {
		res := transaction.Create(&data)
		err = res.Error
	} else {
		res := ctx.DB.Create(&data)
		err = res.Error
	}
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (ctx respositoryUser) GetProductByCode(code string) (*entity.Product, error) {
	var product *entity.Product
	res := ctx.DB.Where(
		&entity.Product{
			ProductCode: code,
		},
	).First(&product)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return product, nil
}
