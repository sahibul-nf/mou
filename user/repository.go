package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method Save() berguna untuk save struct User ke db
func (r *repository) Save(user User) (User, error) {
	// method db.Create() berguna
	// untuk membuat atau menyimpan data baru ke db
	err := r.db.Create(&user).Error

	if err != nil {
		// jika ada error maka kembalikan user dan error
		return user, err
	}

	// jika tidak ada error
	// maka kembalikan user dan nil(tidak ada error)
	return user, nil
}

// method FindEmail() berguna untuk mencari email yang sesuai
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// method FindID() berguna untuk mencari ID yang sesuai
func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// method Update() berguna untuk update avatar user
func (r *repository) Update(user User) (User, error) {
	// method db.Save() berguna
	// untuk membuat perubahan data ke db
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
