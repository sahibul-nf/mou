package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

// TODO: Service RegisterUser
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// mapping input user ke struct User
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	// simpan user baru ke db via repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// TODO: Service LoginUser
func (s *service) LoginUser(input LoginUserInput) (User, error) {
	// dapatkan input user
	email := input.Email
	password := input.PasswordHash

	// cek kesesuaian email
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	// handling jika user input tidak sesuai dengan db
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	// cek user  kesesuaian password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

// TODO: Service Check Available Email
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	// dapatkan input user
	email := input.Email

	// sesuaikan input user berdasarkan email
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	// cek jika email tidak ada yang sesuai
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// TODO: Service Save Avatar / Upload Avatar
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// dapatkan user berdasarkan ID
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// update atrtribute avatar file name
	user.AvatarFileName = fileLocation

	// simpan perubahan avatar file name
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {

	// dapatkan user berdasarkan ID
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// jika user tidak sesuai
	if user.ID == 0 {
		return user, errors.New("No user found on that ID")
	}

	return user, nil
}
