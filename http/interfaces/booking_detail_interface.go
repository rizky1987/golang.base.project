package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type BookingDetailInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.BookingDetail) (*SQLEntity.BookingDetail, error)
	CreateBulk(transaction *gorm.DB, data []*SQLEntity.BookingDetail) error
}
