package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	RoomType struct {
		ID           mssql.UniqueIdentifier `gorm:"column:Id"`
		Code         string                 `gorm:"column:Code"`
		Name         string                 `gorm:"column:Name"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}
)

func (x RoomType) TableName() string {
	return "RoomType"
}

func (u *RoomType) BeforeCreate(tx *gorm.DB) error {
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
