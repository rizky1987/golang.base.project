package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type FloorInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.Floor) (*SQLEntity.Floor, error)
	GetFloorByNumber(number int) (*SQLEntity.Floor, error)
}
