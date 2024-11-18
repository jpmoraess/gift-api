package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/infra/storage"
)

type FileHandler struct {
	FileService *storage.FileService
}

func NewFileHandler(fileService *storage.FileService) *FileHandler {
	return &FileHandler{FileService: fileService}
}

func (fh *FileHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid file"})
	}

	destPath := "uploads/" + file.Filename // Organizar pastas conforme necessário

	fileRecord, err := fh.FileService.Upload(c.Context(), file, destPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fileRecord)
}

func (fh *FileHandler) Download(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	data, err := fh.FileService.Download(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	return c.Send(data)
}

func (fh *FileHandler) Delete(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	if err := fh.FileService.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
