package handlers

import (
	"doko/middleware"
	"doko/models"

	"github.com/gofiber/fiber/v3"
)

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagHandler struct {
	tagRepo *models.TagRepository
}

func NewTagHandler(tagRepo *models.TagRepository) *TagHandler {
	return &TagHandler{
		tagRepo: tagRepo,
	}
}

func (th *TagHandler) GetAllTags(c fiber.Ctx) error {
	tags, err := th.tagRepo.GetAll()
	if err != nil {
		return middleware.NewAppError(fiber.StatusInternalServerError, "Error on fetching tags", err)
	}

	var out []TagResponse
	for _, t := range tags {
		out = append(out, TagResponse{
			ID:   t.ID,
			Name: t.Name,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success fetch tags",
		"data":    out,
	})
}
