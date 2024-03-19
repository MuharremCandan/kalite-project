package service

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	CreateService(user model.User) error
	DeleteService(userID uuid.UUID) error
	UpdateService(user model.User) error
	GetUserService(userID uuid.UUID) (model.User, error)
}
type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo: repo}
}

// Create implements IUserService.
func (u *userService) CreateService(user model.User) error {
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
