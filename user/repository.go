package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method Save() berguna untuk save struct User ke db
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		// jika ada error maka kembalikan user dan error
		return user, err
	}

	// jika tidak ada error
	// maka kembalikan user dan nil(tidak ada error)
	return user, nil
}
