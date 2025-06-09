package services

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/ai/openai"
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/prediction"
)

type PredictionService interface {
	GetPrediction(ctx context.Context, promiseID uuid.UUID) (*dto.PredictionResponse, error)
	CreateOrReplacePrediction(ctx context.Context, promise *models.Promise) (*dto.PredictionResponse, error)
	CreateOrReplacePredictionTx(ctx context.Context, tx *gorm.DB, promise *models.Promise) (*models.Prediction, error)
}

type predictionService struct {
	repo   prediction.PredictionRepository
	client *openai.OpenAIClient
}

func NewPredictionService(repo prediction.PredictionRepository, client *openai.OpenAIClient) PredictionService {
	return &predictionService{
		repo:   repo,
		client: client,
	}
}

func (s *predictionService) GetPrediction(ctx context.Context, promiseID uuid.UUID) (*dto.PredictionResponse, error) {
	existing, err := s.repo.GetByPromiseID(ctx, promiseID)
	if err != nil || existing == nil {
		return nil, err
	}

	return &dto.PredictionResponse{
		PromiseID:   existing.PromiseID.String(),
		SuccessRate: existing.SuccessRate,
		Advice:      existing.Advice,
	}, nil
}

func (s *predictionService) CreateOrReplacePrediction(ctx context.Context, promise *models.Promise) (*dto.PredictionResponse, error) {
	prompt := buildPrompt(promise)

	result, err := s.client.SendPrompt(ctx, prompt)
	if err != nil {
		return nil, err
	}

	successRate, advice := parsePredictionResult(result)

	existing, err := s.repo.GetByPromiseID(ctx, promise.ID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		existing.SuccessRate = float64(successRate)
		existing.Advice = advice
		existing.UpdatedAt = time.Now()
		if err := s.repo.Update(ctx, existing); err != nil {
			return nil, err
		}
		return &dto.PredictionResponse{
			PromiseID:   existing.PromiseID.String(),
			SuccessRate: existing.SuccessRate,
			Advice:      existing.Advice,
		}, nil
	}

	newPrediction := &models.Prediction{
		ID:          uuid.New(),
		PromiseID:   promise.ID,
		SuccessRate: float64(successRate),
		Advice:      advice,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.Create(ctx, newPrediction); err != nil {
		return nil, err
	}

	return &dto.PredictionResponse{
		PromiseID:   newPrediction.PromiseID.String(),
		SuccessRate: newPrediction.SuccessRate,
		Advice:      newPrediction.Advice,
	}, nil
}

func (s *predictionService) CreateOrReplacePredictionTx(ctx context.Context, tx *gorm.DB, promise *models.Promise) (*models.Prediction, error) {
	prompt := buildPrompt(promise)

	result, err := s.client.SendPrompt(ctx, prompt)
	if err != nil {
		return nil, err
	}

	successRate, advice := parsePredictionResult(result)

	if err := s.repo.DeleteByPromiseIDTx(ctx, tx, promise.ID); err != nil {
		return nil, err
	}

	prediction := &models.Prediction{
		ID:          uuid.New(),
		PromiseID:   promise.ID,
		SuccessRate: float64(successRate),
		Advice:      advice,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateTx(ctx, tx, prediction); err != nil {
		return nil, err
	}

	return prediction, nil
}

func buildPrompt(promise *models.Promise) string {
	return "The user created a new personal goal. Here are the details:\n\n" +
		"Title: " + promise.Title + "\n" +
		"Description: " + promise.Description + "\n" +
		"Deadline: " + promise.Deadline.Format("2006-01-02") + "\n" +
		"Today: " + time.Now().Format("2006-01-02") + "\n" +
		"Category: " + promise.Category + "\n\n" +
		"Based on this information, estimate the user's likelihood of success (as a percentage), and give a short motivational advice (max 100 words).\n\n" +
		"Format strictly as:\nSuccess Rate: XX%\nAdvice: ..."
}

func parsePredictionResult(response string) (int, string) {
	lines := strings.Split(response, "\n")
	var successRate int
	var advice string

	for _, line := range lines {
		if strings.HasPrefix(strings.ToLower(line), "success rate:") {
			line = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Success Rate:"), "%"))
			fmt.Sscanf(line, "%d", &successRate)
		}
		if strings.HasPrefix(strings.ToLower(line), "advice:") {
			advice = strings.TrimSpace(strings.TrimPrefix(line, "Advice:"))
		}
	}

	return successRate, advice
}
