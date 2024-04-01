package service

import (
	"errors"
	"fmt"
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateService(user *model.User) error
	DeleteService(userID uuid.UUID) error
	UpdateService(user model.User) error
	GetUserService(userID uuid.UUID) (model.User, error)
	GetUserByUserName(userName string) (model.User, error)
	GetUserByUserMail(userMail string) (model.User, error)
	LoginUserService(identifier, password string) (model.User, error) // identifier = username or user email address
	RegisterUserService(user *model.User) error
}
type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo: repo}
}

// RegisterUserService implements IUserService.
func (u *userService) RegisterUserService(user *model.User) error {

	if err := user.ValidateEmail(); err != nil {
		return err
	}
	if err := user.PassHash(); err != nil {
		return err
	}
	err := u.CreateService(user)
	if err != nil {
		return err
	}
	return nil
}

// LoginUserService implements IUserService.
func (u *userService) LoginUserService(identifier, password string) (model.User, error) {
	// Kullanıcıyı veritabanında bul
	user, err := u.repo.GetUserByUsernameOrMail(identifier)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	// check user password
	if user.ValidateHashPass(user.Password, password) {
		return model.User{}, fmt.Errorf("invalid password")
	}
	return *user, nil
}

// GetUserByUserMail implements IUserService.
func (u *userService) GetUserByUserMail(userMail string) (model.User, error) {
	return u.repo.GetUserByUserMail(userMail)
}

// GetUserByUserName implements IUserService.
func (u *userService) GetUserByUserName(userName string) (model.User, error) {
	return u.repo.GetUserByUserName(userName)
}

// Create implements IUserService.
func (u *userService) CreateService(user *model.User) error {
	user.ID = uuid.New()
	return u.repo.Create(user)
}

// Delete implements IUserService.
func (u *userService) DeleteService(userID uuid.UUID) error {
	return u.repo.Delete(userID)
}

// GetUser implements IUserService.
func (u *userService) GetUserService(userID uuid.UUID) (model.User, error) {
	return u.repo.GetUser(userID)
}

// Update implements IUserService.
func (u *userService) UpdateService(user model.User) error {
	return u.repo.Update(user)
}
