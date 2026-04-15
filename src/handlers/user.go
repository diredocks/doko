package handlers

import (
	"strconv"

	"doko/middleware"
	"doko/models"

	"github.com/gofiber/fiber/v3"
)

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterUserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type DeleteUserRequest struct {
	ID uint `params:"id" validate:"required"`
}

type UserHandler struct {
	userRepo *models.UserRepository
}

func NewUserHandler(userRepo *models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (uh *UserHandler) Delete(c fiber.Ctx) error {
	var in DeleteUserRequest
	if err := c.Bind().URI(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	if err := uh.userRepo.Delete(in.ID); err != nil {
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on deleting user", err)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success delete user",
		"data":    nil,
	})
}

func (uh *UserHandler) Register(c fiber.Ctx) error {
	var in RegisterUserRequest
	if err := c.Bind().Body(&in); err != nil {
		return middleware.NewValidationError(err)
	}

	user, err := uh.userRepo.Create(in.Email, in.Username, in.Password)
	if err != nil {
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on registering user", err)
	}

	newUser := RegisterUserResponse{
		ID:       strconv.FormatUint(uint64(user.ID), 10),
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success register",
		"data":    newUser,
	})
}
