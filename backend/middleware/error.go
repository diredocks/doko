package middleware

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

// AppError is a domain-specific error with public and internal details.
type AppError struct {
	Code      int
	PublicMsg string
	Internal  error
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return e.Internal.Error()
	}
	return e.PublicMsg
}

// ValidationError carries structured per-field validation failures.
type ValidationError struct {
	Fields map[string]string
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

// NewAppError creates a new AppError.
func NewAppError(code int, publicMsg string, internal error) *AppError {
	return &AppError{Code: code, PublicMsg: publicMsg, Internal: internal}
}

// NewValidationError creates a ValidationError from go-playground validator errors.
func NewValidationError(err error) *ValidationError {
	ve := &ValidationError{Fields: make(map[string]string)}
	if valErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
		for _, fe := range valErrs {
			ve.Fields[fe.Field()] = fe.Tag()
		}
	} else {
		ve.Fields["general"] = err.Error()
	}
	return ve
}

// ErrorHandler is the central Fiber error handler.
func ErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	publicMsg := "Internal Server Error"

	// 1) Structured validation errors (422)
	if valErr, ok := errors.AsType[*ValidationError](err); ok {
		return respond(c, fiber.StatusUnprocessableEntity, "Validation failed", fiber.Map{
			"errors": valErr.Fields,
		})
	}

	// 2) Custom domain errors
	if appErr, ok := errors.AsType[*AppError](err); ok {
		code = appErr.Code
		if code == 0 {
			code = fiber.StatusInternalServerError
		}
		publicMsg = appErr.PublicMsg
		if publicMsg == "" {
			publicMsg = "Internal Server Error"
		}
		log.Printf("[AppError] %d %s: %v", code, publicMsg, appErr.Internal)
		return respond(c, code, publicMsg, nil)
	}

	// 3) Fiber *fiber.Error
	if fiberErr, ok := errors.AsType[*fiber.Error](err); ok {
		code = fiberErr.Code
		publicMsg = fiberErr.Message
		log.Printf("[FiberError] %d: %s", code, fiberErr.Message)
		return respond(c, code, publicMsg, nil)
	}

	// 4) Fallback — log the real error, never leak it
	log.Printf("[UnhandledError] %v", err)
	return respond(c, code, publicMsg, nil)
}

func respond(c fiber.Ctx, code int, message string, data any) error {
	accept := c.Get("Accept")
	if strings.Contains(accept, "text/html") {
		c.Status(code)
		return c.SendString(message)
	}

	payload := fiber.Map{
		"status":  "error",
		"message": message,
	}
	if data != nil {
		payload["data"] = data
	}
	c.Status(code)
	return c.JSON(payload)
}

func Recover() fiber.Handler {
	return recoverer.New()
}
