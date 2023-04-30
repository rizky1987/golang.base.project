package entities

import (
	"errors"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	BookingDetail struct {
		ID        mssql.UniqueIdentifier `gorm:"column:Id"`
		BookingId mssql.UniqueIdentifier `gorm:"column:BookingId"`
		RoomId    mssql.UniqueIdentifier `gorm:"column:RoomId"`
		Price     int                    `gorm:"column:Price"`
	}
)

func (x BookingDetail) TableName() string {
	return "BookingDetail"
}

func (u *BookingDetail) BeforeCreate(tx *gorm.DB) error {
	err := u.ID.Scan(u.ID.String())

	if err != nil || u.ID.String() == "00000000-0000-0000-0000-000000000000" { // if Id is null or not valid UUID

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
	}

	return nil
}
