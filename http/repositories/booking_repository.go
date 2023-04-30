package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"
	"strings"

	"gorm.io/gorm"
)

type respositoryBooking struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) interfaces.BookingInterface {
	return &respositoryBooking{db}
}

func (ctx respositoryBooking) Create(transaction *gorm.DB, data *SQLEntity.Booking) error {
	var err error

	if transaction != nil {
		res := transaction.Create(&data)
		err = res.Error
	} else {
		res := ctx.DB.Create(&data)
		err = res.Error
	}
	if err != nil {
		return err
	}
	return nil
}

func (ctx respositoryBooking) GetBookingByCode(code string) (*SQLEntity.Booking, error) {
	var room *SQLEntity.Booking
	res := ctx.DB.Where(
		&SQLEntity.Booking{
			Code: code,
		},
	).First(&room)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return room, nil
}

func (ctx respositoryBooking) GetBookingRoomAvaibility(roomIds []string, startDate, endDate string) ([]*SQLEntity.TempBookingRoomAvaibility, error) {
	var err error
	var results []*SQLEntity.TempBookingRoomAvaibility

	err = ctx.DB.Raw(`SELECT 
						b.Code as BookingCode,
						b.BookedName as BookingName,
						b.StartDate as BookingStartDate,
						b.EndDate as BookingEndDate,
						r.Code as RoomCode
					FROM Booking b
						left join BookingDetail bd on b.Id = bd.BookingId
						left join Room r on r.Id = bd.RoomId
					WHERE ((CONVERT(DATE,StartDate) <= ? and CONVERT(DATE,EndDate) >= ?)
						or (CONVERT(DATE,StartDate) <= ? and CONVERT(DATE,EndDate) >= ?))
						AND bd.RoomId in (?)`, startDate, startDate, endDate, endDate, strings.Join(roomIds, ",")).
		Find(&results).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return results, nil
}
