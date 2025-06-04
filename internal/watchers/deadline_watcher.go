package watchers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/services"
)

var deadlineDays = []int{1, 2, 3, 4, 5, 6, 7}
var deadlineMonths = []int{1, 2, 3, 6, 12}

func StartDeadlineWatcher(
	ctx context.Context,
	promiseRepo promise.PromiseRepository,
	notificationService services.NotificationService,
) {
	go func() {
		for {
			now := time.Now()
			nextRun := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			time.Sleep(time.Until(nextRun))

			log.Println("🔔 Running deadline notification watcher...")
			promises, err := promiseRepo.GetAllPromises(ctx)
			if err != nil {
				log.Println("❌ Failed to get promises:", err)
				continue
			}

			for _, p := range promises {
				sendDeadlineReminderIfNeeded(ctx, &p, notificationService)
			}
		}
	}()
}

func sendDeadlineReminderIfNeeded(ctx context.Context, promise *models.Promise, notificationService services.NotificationService) {
	if promise.Status == "completed" {
		return
	}

	diff := promise.Deadline.Sub(time.Now())
	days := int(diff.Hours() / 24)

	// check days
	for _, day := range deadlineDays {
		if days == day {
			sendNotification(ctx, promise.UserID, promise.ID, notificationService, day, "day")
			return
		}
	}

	// check months
	for _, m := range deadlineMonths {
		targetDate := promise.Deadline.AddDate(0, -m, 0)
		if time.Now().Year() == targetDate.Year() && time.Now().YearDay() == targetDate.YearDay() {
			sendNotification(ctx, promise.UserID, promise.ID, notificationService, m, "month")
			break
		}
	}
}

func sendNotification(
	ctx context.Context,
	userID uuid.UUID,
	promiseID uuid.UUID,
	notificationService services.NotificationService,
	amount int,
	unit string,
) {
	var msg string

	switch unit {
	case "day":
		if amount == 1 {
			msg = "⏰ Tomorrow is the deadline for one of your promises!"
		} else {
			msg = fmt.Sprintf("⏰ %d days left until your promise deadline.", amount)
		}
	case "month":
		if amount == 12 {
			msg = "⏰ 1 year remains until your promise deadline."
		} else {
			msg = fmt.Sprintf("⏰ %d months left until your promise deadline.", amount)
		}
	default:
		msg = "⏰ A reminder about your promise deadline."
	}

	notif := &models.Notification{
		Type:      "deadline",
		Message:   msg,
		RelatedID: &promiseID,
		CreatedAt: time.Now(),
	}

	_ = notificationService.SendNotification(ctx, userID, notif)
}
