package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type RoomInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.Room) (*SQLEntity.Room, error)
	GetRoomByCode(code string) (*SQLEntity.Room, error)
}
