package handlers

import (
	"fmt"
	"strconv"

	"doko/models"

	"github.com/gofiber/fiber/v3"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
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

	user, err := uh.userRepo.CreateUser(in.Email, in.Username, in.Password)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	newUser := RegisterResponse{
		Id:       strconv.FormatUint(uint64(user.ID), 10),
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success register",
		"data":    newUser,
	})
}
