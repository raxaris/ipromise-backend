package attachment

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

func (r *attachmentRepository) UploadAttachment(ctx context.Context, attachment *models.Attachment) error {
	return r.db.WithContext(ctx).Create(attachment).Error
}

func (r *attachmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Attachment, error) {
	var attachment models.Attachment
	if err := r.db.WithContext(ctx).First(&attachment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *attachmentRepository) ListAttachmentsByPostID(
	ctx context.Context,
	postID uuid.UUID,
) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.WithContext(ctx).
		Where("post_id = ?", postID).
		Find(&attachments).Error
	return attachments, err
}

func (r *attachmentRepository) DeleteAttachment(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Delete(&models.Attachment{}, "id = ?", id).Error
}

func (r *attachmentRepository) DeleteAllByPostID(ctx context.Context, postID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("post_id = ?", postID).
		Delete(&models.Attachment{}).Error
}
