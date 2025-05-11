package test

import (
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

func LoadTestPostTree(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Основной пост
		post1 := models.Post{
			ID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			UserID:      uuid.MustParse("083b262d-6d16-4b39-8b1e-34f2b4d7e682"),
			PromiseID:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			MicrotaskID: uuid.MustParse("00000000-0000-0000-0000-000000000101"),
			Content:     "Я купил машину",
			CreatedAt:   time.Now(),
		}

		// Первый комментарий
		post2 := models.Post{
			ID:          uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			UserID:      uuid.MustParse("7014a87b-164b-40f8-b264-c73117271765"),
			PromiseID:   post1.PromiseID,
			MicrotaskID: uuid.MustParse("00000000-0000-0000-0000-000000000102"),
			Content:     "Ого, сколько стоит?",
			ParentID:    &post1.ID,
			CreatedAt:   time.Now(),
		}

		// Ответ на комментарий
		post3 := models.Post{
			ID:          uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			UserID:      uuid.MustParse("812596ac-cbe0-40ad-ade2-b70edb0f3b75"),
			PromiseID:   post1.PromiseID,
			MicrotaskID: uuid.MustParse("00000000-0000-0000-0000-000000000103"),
			Content:     "18 лямов",
			ParentID:    &post2.ID,
			CreatedAt:   time.Now(),
		}

		// Второй комментарий
		post4 := models.Post{
			ID:          uuid.MustParse("44444444-4444-4444-4444-444444444444"),
			UserID:      uuid.MustParse("9036bfdd-fcc6-4993-930c-e1f7fdaa71cf"),
			PromiseID:   post1.PromiseID,
			MicrotaskID: uuid.MustParse("00000000-0000-0000-0000-000000000104"),
			Content:     "Поздравляю!",
			ParentID:    &post1.ID,
			CreatedAt:   time.Now(),
		}

		// Ответ на второй комментарий
		post5 := models.Post{
			ID:          uuid.MustParse("55555555-5555-5555-5555-555555555555"),
			UserID:      uuid.MustParse("b09b93fd-5c27-434f-a953-e4fe172129d0"),
			PromiseID:   post1.PromiseID,
			MicrotaskID: uuid.MustParse("00000000-0000-0000-0000-000000000105"),
			Content:     "Спасибо!",
			ParentID:    &post4.ID,
			CreatedAt:   time.Now(),
		}

		posts := []*models.Post{&post1, &post2, &post3, &post4, &post5}

		for _, p := range posts {
			if err := tx.Create(p).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
