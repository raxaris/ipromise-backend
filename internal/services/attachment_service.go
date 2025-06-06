package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/attachment"
	"github.com/raxaris/ipromise-backend/internal/storage"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

type AttachmentService interface {
	UploadAttachment(ctx context.Context, postID uuid.UUID, userID uuid.UUID, fileBytes []byte, fileName, fileType string) (*models.Attachment, error)
	UploadAttachmentToPost(ctx context.Context, postID uuid.UUID, fileBytes []byte, fileName, fileType string) (*models.Attachment, error)
	UploadAvatar(ctx context.Context, userID uuid.UUID, fileBytes []byte, fileName, fileType string) (*models.Attachment, error)
	ListAttachmentsByPostID(ctx context.Context, postID uuid.UUID) ([]models.Attachment, error)
	DeleteAttachment(ctx context.Context, id uuid.UUID) error
	DeleteAllByPostID(ctx context.Context, postID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Attachment, error)
}

type attachmentService struct {
	repo    attachment.AttachmentRepository
	storage storage.Storage
}

func NewAttachmentService(
	repo attachment.AttachmentRepository,
	storage storage.Storage,
) AttachmentService {
	return &attachmentService{
		repo:    repo,
		storage: storage,
	}
}

func (s *attachmentService) UploadAttachment(
	ctx context.Context,
	postID, userID uuid.UUID,
	fileBytes []byte,
	fileName, fileType string,
) (*models.Attachment, error) {
	fileURL, err := s.storage.UploadFile(ctx, fileBytes, fileName)
	if err != nil {
		return nil, err
	}

	var attachmentType string
	var postPtr *uuid.UUID
	var userPtr *uuid.UUID

	if postID != uuid.Nil {
		attachmentType = "post_image"
		postPtr = &postID
	}
	if userID != uuid.Nil {
		attachmentType = "avatar"
		userPtr = &userID
	}

	if postPtr == nil && userPtr == nil {
		_ = s.storage.DeleteFile(ctx, fileURL)
		return nil, utils.ErrInvalidInput
	}

	newAttachment := &models.Attachment{
		ID:             uuid.New(),
		PostID:         postPtr,
		UserID:         userPtr,
		AttachmentType: attachmentType,
		FileURL:        fileURL,
		FileType:       fileType,
	}

	if err := s.repo.UploadAttachment(ctx, newAttachment); err != nil {
		_ = s.storage.DeleteFile(ctx, fileURL)
		return nil, err
	}

	return newAttachment, nil
}

func (s *attachmentService) UploadAttachmentToPost(
	ctx context.Context,
	postID uuid.UUID,
	fileBytes []byte,
	fileName, fileType string,
) (*models.Attachment, error) {
	if postID == uuid.Nil {
		return nil, utils.ErrInvalidInput
	}

	fileURL, err := s.storage.UploadFile(ctx, fileBytes, fileName)
	if err != nil {
		return nil, err
	}

	newAttachment := &models.Attachment{
		ID:             uuid.New(),
		PostID:         &postID,
		AttachmentType: "post_image",
		FileURL:        fileURL,
		FileType:       fileType,
	}

	if err := s.repo.UploadAttachment(ctx, newAttachment); err != nil {
		_ = s.storage.DeleteFile(ctx, fileURL)
		return nil, err
	}

	return newAttachment, nil
}

func (s *attachmentService) UploadAvatar(
	ctx context.Context,
	userID uuid.UUID,
	fileBytes []byte,
	fileName, fileType string,
) (*models.Attachment, error) {
	if userID == uuid.Nil {
		return nil, utils.ErrInvalidInput
	}

	fileURL, err := s.storage.UploadFile(ctx, fileBytes, fileName)
	if err != nil {
		return nil, err
	}

	newAttachment := &models.Attachment{
		ID:             uuid.New(),
		UserID:         &userID,
		AttachmentType: "avatar",
		FileURL:        fileURL,
		FileType:       fileType,
	}

	if err := s.repo.UploadAttachment(ctx, newAttachment); err != nil {
		_ = s.storage.DeleteFile(ctx, fileURL)
		return nil, err
	}

	return newAttachment, nil
}

func (s *attachmentService) ListAttachmentsByPostID(ctx context.Context, postID uuid.UUID) ([]models.Attachment, error) {
	return s.repo.ListAttachmentsByPostID(ctx, postID)
}

func (s *attachmentService) DeleteAttachment(ctx context.Context, id uuid.UUID) error {
	existingAttachment, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteAttachment(ctx, id); err != nil {
		return err
	}

	return s.storage.DeleteFile(ctx, existingAttachment.FileURL)
}

func (s *attachmentService) DeleteAllByPostID(ctx context.Context, postID uuid.UUID) error {
	attachments, err := s.repo.ListAttachmentsByPostID(ctx, postID)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteAllByPostID(ctx, postID); err != nil {
		return err
	}

	for _, att := range attachments {
		_ = s.storage.DeleteFile(ctx, att.FileURL)
	}

	return nil
}

func (s *attachmentService) GetByID(ctx context.Context, id uuid.UUID) (*models.Attachment, error) {
	return s.repo.GetByID(ctx, id)
}
