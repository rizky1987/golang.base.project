package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type RoomTypeInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.RoomType) (*SQLEntity.RoomType, error)
	GetRoomTypeByCode(code string) (*SQLEntity.RoomType, error)
}
