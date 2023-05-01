package interfaces

import (
	SQLEntity "hotel/databases/entities/sql"

	"gorm.io/gorm"
)

type UserInterface interface {
	Create(transaction *gorm.DB, data SQLEntity.User) (*SQLEntity.User, error)
	GetUserByUsername(username string) (*SQLEntity.User, error)
	GetUserByUsernamePassword(username, password string) (*SQLEntity.User, error)
}
