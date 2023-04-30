package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"
	"strings"

	"gorm.io/gorm"
)

type respositoryRoom struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) interfaces.RoomInterface {
	return &respositoryRoom{db}
}

func (ctx respositoryRoom) Create(transaction *gorm.DB, data SQLEntity.Room) (*SQLEntity.Room, error) {
	var err error

	if transaction != nil {
		res := transaction.Create(&data)
		err = res.Error
	} else {
		res := ctx.DB.Create(&data)
		err = res.Error
	}
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (ctx respositoryRoom) GetRoomByCode(code string) (*SQLEntity.Room, error) {
	var room *SQLEntity.Room
	res := ctx.DB.Where(
		&SQLEntity.Room{
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

func (ctx respositoryRoom) GetRoomDetailByRoomIds(roomIds []string) ([]*SQLEntity.TempRoomDetail, error) {
	var err error
	var results []*SQLEntity.TempRoomDetail

	err = ctx.DB.Raw(`SELECT 
						r.Id as RoomId,
						r.Code as RoomCode,
						r.Number as RoomNumber,
						f.Id as FloorId,
						f.Number as FloorNumber,
						f.Price as FloorPrice,
						rt.Id as RoomTypeId,
						rt.Code as RoomTypeCode,
						rt.[Name] as RoomTypeName 
					FROM Room r
						left join [Floor] f on r.FloorId = f.Id
						left join [RoomType] rt on f.RoomTypeId = rt.Id
						WHERE r.Id in (?) `, strings.Join(roomIds, ",")).
		Find(&results).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return results, nil
}
