package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type RoomPriceInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.RoomPrice) (*SQLEntity.RoomPrice, error)
	GetRoomPriceByCode(code string) (*SQLEntity.RoomPrice, error)
}
