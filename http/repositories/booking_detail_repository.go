package repositories

import (
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"

	"gorm.io/gorm"
)

type respositoryBookingDetail struct {
	DB *gorm.DB
}

func NewBookingDetailRepository(db *gorm.DB) interfaces.BookingDetailInterface {
	return &respositoryBookingDetail{db}
}

func (ctx respositoryBookingDetail) Create(transaction *gorm.DB, data SQLEntity.BookingDetail) (*SQLEntity.BookingDetail, error) {
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

func (ctx respositoryBookingDetail) CreateBulk(transaction *gorm.DB, data []*SQLEntity.BookingDetail) error {
	var err error

	if transaction != nil {
		res := transaction.CreateInBatches(&data, 100)
		err = res.Error
	} else {
		res := ctx.DB.CreateInBatches(&data, 100)
		err = res.Error
	}

	if err != nil {
		return err
	}

	return nil
}
