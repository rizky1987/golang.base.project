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
		RoomPriceId  mssql.UniqueIdentifier `gorm:"column:RoomPriceId"`
		Code         string                 `gorm:"column:Code"`
		Number       int                    `gorm:"column:Number"`
		CreatedBy    string                 `gorm:"column:CreatedBy"`
		CreatedDate  time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy   *string                `gorm:"column:ModifiedBy"`
		ModifiedDate *time.Time             `gorm:"column:ModifiedDate"`
		DeletedBy    *string                `gorm:"column:DeletedBy"`
		DeletedDate  *time.Time             `gorm:"column:DeletedDate"`
	}

	TempRoomDetail struct {
		RoomId         mssql.UniqueIdentifier `gorm:"column:RoomId"`
		RoomCode       string                 `gorm:"column:RoomCode"`
		RoomNumber     int                    `gorm:"column:RoomNumber"`
		FloorId        mssql.UniqueIdentifier `gorm:"column:FloorId"`
		RoomPricePrice int                    `gorm:"column:RoomPricePrice"`
		RoomTypeId     mssql.UniqueIdentifier `gorm:"column:RoomTypeId"`
		RoomTypeCode   string                 `gorm:"column:RoomTypeCode"`
		RoomTypeName   string                 `gorm:"column:RoomTypeName"`
	}

	TempAvaibilityRoom struct {
		FloorNumber    int    `gorm:"column:FloorNumber"`
		RoomNumber     int    `gorm:"column:RoomNumber"`
		RoomPriceType  string `gorm:"column:RoomPriceType"`
		RoomPricePrice int    `gorm:"column:RoomPricePrice"`
		IsBooked       int    `gorm:"column:IsBooked"`
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
