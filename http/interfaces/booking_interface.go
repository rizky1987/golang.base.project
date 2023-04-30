package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type BookingInterface interface {
	Create(transaction *gorm.DB, data *SQLEntity.Booking) error
	GetBookingByCode(code string) (*SQLEntity.Booking, error)
	GetBookingRoomAvaibility(roomIds []string, startDate, endDate string) ([]*SQLEntity.TempBookingRoomAvaibility, error)
}
