package repositories

import (
	"errors"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"

	"gorm.io/gorm"
)

type respositoryUser struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserInterface {
	return &respositoryUser{db}
}

func (ctx respositoryUser) Create(transaction *gorm.DB, data SQLEntity.User) (*SQLEntity.User, error) {
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

func (ctx respositoryUser) GetUserByUsername(username string) (*SQLEntity.User, error) {
	var user *SQLEntity.User
	res := ctx.DB.Where(
		&SQLEntity.User{
			Username: username,
		},
	).First(&user)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

func (ctx respositoryUser) GetUserByUsernamePassword(username, password string) (*SQLEntity.User, error) {
	var user *SQLEntity.User
	res := ctx.DB.Where(
		&SQLEntity.User{
			Username: username,
			Password: password,
		},
	).First(&user)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}
