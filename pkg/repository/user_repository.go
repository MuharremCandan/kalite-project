package repository

import (
	"go-backend-test/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) error
	Delete(userID uuid.UUID) error
	Update(user model.User) error
	GetUser(userID uuid.UUID) (model.User, error)
	GetUserByUserName(userName string) (model.User, error)
	GetUserByUserMail(userMail string) (model.User, error)
	GetUserByUsernameOrMail(usernameOrMail string) (*model.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// GetUserByUsernameOrMail implements IUserRepository.
func (u *userRepository) GetUserByUsernameOrMail(usernameOrMail string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("user_name = ? OR email = ?", usernameOrMail, usernameOrMail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUserMail implements IUserRepository.
func (u *userRepository) GetUserByUserMail(userMail string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", userMail).First(&user).Error
	return user, err
}

// GetUserByUserName implements IUserRepository.
func (u *userRepository) GetUserByUserName(userName string) (model.User, error) {
	var user model.User
	err := u.db.Where("user_name = ?", userName).First(&user).Error
	return user, err
}

// Create implements IUserRepository.
func (u *userRepository) Create(user *model.User) error {
	return u.db.Create(user).Error
}

// Delete implements IUserRepository.
func (u *userRepository) Delete(userID uuid.UUID) error {
	return u.db.Where("id = ?", userID).Delete(&model.User{}).Error
}

// GetUser implements IUserRepository.
func (u *userRepository) GetUser(userID uuid.UUID) (model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

// Update implements IUserRepository.
func (u *userRepository) Update(user model.User) error {
	return u.db.Save(&user).Error
}
