package attachment

import (
	"context"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
)

type AttachmentRepository interface {
	UploadAttachment(ctx context.Context, attachment *models.Attachment) error
	ListAttachmentsByPostID(ctx context.Context, postID uuid.UUID) ([]models.Attachment, error)
	DeleteAttachment(ctx context.Context, id uuid.UUID) error
}
