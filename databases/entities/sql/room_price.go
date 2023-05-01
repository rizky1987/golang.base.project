package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	RoomPrice struct {
		ID           mssql.UniqueIdentifier `gorm:"column:Id"`
		Code         string                 `gorm:"column:Code"`
		Type         string                 `gorm:"column:Type"`
		Price        string                 `gorm:"column:Price"`
		FloorId      mssql.UniqueIdentifier `gorm:"column:FloorId"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}
)

func (x RoomPrice) TableName() string {
	return "RoomPrice"
}

func (u *RoomPrice) BeforeCreate(tx *gorm.DB) error {
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
