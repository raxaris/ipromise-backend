package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/services"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
)

type AttachmentHandler struct {
	attachmentService services.AttachmentService
}

func NewAttachmentHandler(attachmentService services.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{attachmentService: attachmentService}
}

// UploadAttachmentsToPost godoc
// @Summary Загрузить вложения к посту
// @Description Загружает файлы, сохраняет в MinIO и БД, возвращает массив с URL-ами
// @Tags attachments
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID поста"
// @Param attachments formData file true "Файлы (можно несколько)"
// @Success 201 {array} dto.AttachmentResponse
// @Failure 400 {object} map[string]string "error: Неверный ID или формат"
// @Failure 401 {object} map[string]string "error: Неавторизован"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /posts/{id}/attachments [post]
func (h *AttachmentHandler) UploadAttachmentsToPost(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID поста")
		return
	}

	form, err := c.MultipartForm()
	if err != nil || form.File == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Нет файлов для загрузки")
		return
	}
	fmt.Println("form", form)
	files := form.File["attachments"]
	var responses []dto.AttachmentResponse
	fmt.Println("files", files)
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("upload error:", err) // 👈 добавь это!
			continue
		}
		defer file.Close()

		fileBytes := make([]byte, fileHeader.Size)
		_, _ = file.Read(fileBytes)

		uploaded, err := h.attachmentService.UploadAttachmentToPost(
			c.Request.Context(),
			postID,
			fileBytes,
			fileHeader.Filename,
			fileHeader.Header.Get("Content-Type"),
		)
		if err != nil {
			fmt.Println("upload error:", err)
			continue
		}

		fmt.Println("uploaded", uploaded)

		responses = append(responses, dto.AttachmentResponse{
			ID:       uploaded.ID.String(),
			FileURL:  uploaded.FileURL,
			FileType: uploaded.FileType,
		})
		fmt.Println("responses", responses)
	}

	utils.RespondWithSuccess(c, http.StatusCreated, responses)
}

// ListAttachmentsByPostID godoc
// @Summary Получить вложения поста
// @Description Возвращает список вложенных файлов у поста
// @Tags attachments
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID поста"
// @Success 200 {array} dto.AttachmentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{id}/attachments [get]
func (h *AttachmentHandler) ListAttachmentsByPostID(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID поста")
		return
	}

	attachments, err := h.attachmentService.ListAttachmentsByPostID(c.Request.Context(), postID)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	var response []dto.AttachmentResponse
	for _, a := range attachments {
		response = append(response, dto.AttachmentResponse{
			ID:       a.ID.String(),
			FileURL:  a.FileURL,
			FileType: a.FileType,
		})
	}

	utils.RespondWithSuccess(c, http.StatusOK, response)
}

// DeleteAttachmentByID godoc
// @Summary Удалить вложение по ID
// @Description Удаляет вложение из MinIO и базы данных
// @Tags attachments
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID вложения"
// @Success 200 {object} map[string]string "message: Удалено"
// @Failure 400 {object} map[string]string "error: Неверный ID"
// @Failure 500 {object} map[string]string "error: Ошибка сервера"
// @Router /attachments/{id} [delete]
func (h *AttachmentHandler) DeleteAttachmentByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Некорректный ID")
		return
	}

	err = h.attachmentService.DeleteAttachment(c.Request.Context(), id)
	if err != nil {
		utils.RespondWithMappedError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Вложение удалено"})
}
