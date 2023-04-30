package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"

	"gorm.io/gorm"
)

type respositoryFloor struct {
	DB *gorm.DB
}

func NewFloorRepository(db *gorm.DB) interfaces.FloorInterface {
	return &respositoryFloor{db}
}

func (ctx respositoryFloor) Create(transaction *gorm.DB, data SQLEntity.Floor) (*SQLEntity.Floor, error) {
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

func (ctx respositoryFloor) GetFloorByNumber(number int) (*SQLEntity.Floor, error) {
	var floor *SQLEntity.Floor
	res := ctx.DB.Where(
		&SQLEntity.Floor{
			Number: number,
		},
	).First(&floor)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return floor, nil
}
