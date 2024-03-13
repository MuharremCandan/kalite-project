package model

import (
	"go-backend-test/pkg/utils"
	"net/mail"

	"github.com/google/uuid"
)

// User struct represents user information
type User struct {
	Base
	Email    string `json:"email" example:"john.doe@example.com"` // Example email
	Name     string `json:"name" example:"John Doe"`              // Example name
	Surname  string `json:"surname" example:"Doe"`                // Example surname
	Phone    string `json:"phone" example:"+1234567890"`          // Example phone (anonymized)
	Address  string `json:"address" example:"123 Main Street"`    // Example address
	Password string `json:"password"`
}

func (u *User) ValidateEmail() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) PassHash() error {
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pass
	return nil
}

// UpdatePasswordRequest struct represents the request payload for updating user password.
type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

// LoginRequest struct represents the request payload for login in
type LoginRequest struct {
	UsernameOrMail string `json:"username_or_mail" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

// RegisterRequest struct represents the request payload for registering
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// LoginResponse struct represents the response payload for login in
type LoginResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"name"`
	Mail     string    `json:"mail"`
}
