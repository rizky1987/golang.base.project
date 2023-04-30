package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	Room struct {
		ID           mssql.UniqueIdentifier `gorm:"column:Id"`
		FloorId      mssql.UniqueIdentifier `gorm:"column:FloorId"`
		Code         string                 `gorm:"column:Code"`
		Number       int                    `gorm:"column:Number"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}
)

func (x Room) TableName() string {
	return "Room"
}

func (u *Room) BeforeCreate(tx *gorm.DB) error {
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
