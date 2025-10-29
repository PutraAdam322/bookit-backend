package repository

import (
	"fmt"

	"bookit.com/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (usr *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	tx := usr.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (usr *UserRepository) GetByIDPreloadBooking(id uint) (*model.User, error) {
	var user model.User
	tx := usr.db.Preload("Bookings").Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (usr *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	tx := usr.db.Where("email = ?", email).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func (usr *UserRepository) Create(user *model.User) error {
	fmt.Println(&user)
	tx := usr.db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (usr *UserRepository) Update(user *model.User) error {
	tx := usr.db.Where("id = ?", user.ID).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
