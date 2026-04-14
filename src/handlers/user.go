package handlers

import (
	"strconv"

	"doko/models"

	"github.com/gofiber/fiber/v3"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserHandler struct {
	userRepo *models.UserRepository
}

func NewUserHandler(userRepo *models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (uh *UserHandler) Register(c fiber.Ctx) error {
	var in RegisterRequest
	if err := c.Bind().Body(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on register request",
			"data":    nil,
		})
	}

	user, err := uh.userRepo.Create(in.Email, in.Username, in.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on registering user",
			"data":    nil,
		})
	}

	newUser := RegisterResponse{
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
