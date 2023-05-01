package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	Booking struct {
		ID           mssql.UniqueIdentifier `gorm:"column:Id"`
		Code         string                 `gorm:"column:Code"`
		StartDate    time.Time              `gorm:"column:StartDate"`
		EndDate      time.Time              `gorm:"column:EndDate"`
		DownPayment  int                    `gorm:"column:DownPayment"`
		IsPaidOff    bool                   `gorm:"column:IsPaidOff"`
		BookedBy     string                 `gorm:"column:BookedBy"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}

	TempBookingRoomAvaibility struct {
		BookingCode      string    `gorm:"column:BookingCode"`
		BookedBy         string    `gorm:"column:BookedBy"`
		BookingStartDate time.Time `gorm:"column:BookingStartDate"`
		BookingEndDate   time.Time `gorm:"column:BookingEndDate"`
		RoomCode         string    `gorm:"column:RoomCode"`
	}
)

func (x Booking) TableName() string {
	return "Booking"
}

func (u *Booking) BeforeCreate(tx *gorm.DB) error {
	err := u.ID.Scan(u.ID.String())

	if err != nil { // if Id is null or not valid UUID
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
