package repository

import (
	"go-backend-test/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) error
	Delete(userID uuid.UUID) error
	Update(user model.User) error
	GetUser(userID uuid.UUID) (model.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// Create implements IUserRepository.
func (u *userRepository) Create(user model.User) error {
	return u.db.Create(user).Error
}

// Delete implements IUserRepository.
func (u *userRepository) Delete(userID uuid.UUID) error {
	return u.db.Where("user_id = ?", userID).Delete(&model.User{}).Error
}

// GetUser implements IUserRepository.
func (u *userRepository) GetUser(userID uuid.UUID) (model.User, error) {
	var user model.User
	err := u.db.Where("user_id = ?", userID).First(&user).Error
	return user, err
}

// Update implements IUserRepository.
func (u *userRepository) Update(user model.User) error {
	return u.db.Save(&user).Error
}
