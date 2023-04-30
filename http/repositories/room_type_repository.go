package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"

	"gorm.io/gorm"
)

type respositoryRoomType struct {
	DB *gorm.DB
}

func NewRoomTypeRepository(db *gorm.DB) interfaces.RoomTypeInterface {
	return &respositoryRoomType{db}
}

func (ctx respositoryRoomType) Create(transaction *gorm.DB, data SQLEntity.RoomType) (*SQLEntity.RoomType, error) {
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

func (ctx respositoryRoomType) GetRoomTypeByCode(code string) (*SQLEntity.RoomType, error) {
	var roomType *SQLEntity.RoomType
	res := ctx.DB.Where(
		&SQLEntity.RoomType{
			Code: code,
		},
	).First(&roomType)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return roomType, nil
}
