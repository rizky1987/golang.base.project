package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"

	"gorm.io/gorm"
)

type respositoryRoomPrice struct {
	DB *gorm.DB
}

func NewRoomPriceRepository(db *gorm.DB) interfaces.RoomPriceInterface {
	return &respositoryRoomPrice{db}
}

func (ctx respositoryRoomPrice) Create(transaction *gorm.DB, data SQLEntity.RoomPrice) (*SQLEntity.RoomPrice, error) {
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

func (ctx respositoryRoomPrice) GetRoomPriceByCode(code string) (*SQLEntity.RoomPrice, error) {
	var roomPrice *SQLEntity.RoomPrice
	res := ctx.DB.Where(
		&SQLEntity.RoomPrice{
			Code: code,
		},
	).First(&roomPrice)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return roomPrice, nil
}
