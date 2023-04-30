package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	Floor struct {
		ID           mssql.UniqueIdentifier `gorm:"column:Id"`
		RoomTypeId   mssql.UniqueIdentifier `gorm:"column:RoomTypeId"`
		Number       int                    `gorm:"column:Number"`
		Price        int                    `gorm:"column:Price"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}
)

func (x Floor) TableName() string {
	return "Floor"
}

func (u *Floor) BeforeCreate(tx *gorm.DB) error {
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
