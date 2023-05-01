package entities

import (
	"errors"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID       mssql.UniqueIdentifier `gorm:"column:Id"`
		Username string                 `gorm:"column:Username"`
		Password string                 `gorm:"column:Password"`
	}
)

func (x User) TableName() string {
	return "User"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u2, err := uuid.NewV4()
	if err != nil {
		errors.New("can't save invalid data")
		return err
	}
	err = u.ID.Scan(u2.String())
	if err != nil {
		errors.New("can't save invalid data")
		return err
	}
	return nil
}
