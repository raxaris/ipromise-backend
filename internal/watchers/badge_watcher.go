package watchers

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/repositories/follower"
	"github.com/raxaris/ipromise-backend/internal/repositories/post"
	"github.com/raxaris/ipromise-backend/internal/repositories/promise"
	"github.com/raxaris/ipromise-backend/internal/repositories/user"
	"github.com/raxaris/ipromise-backend/internal/services"
)

func StartBadgeWatcher(
	ctx context.Context,
	userRepo user.UserRepository,
	postRepo post.PostRepository,
	promiseRepo promise.PromiseRepository,
	followerRepo follower.FollowerRepository,
	badgeService services.BadgeService,
) {
	ticker := time.NewTicker(60 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				users, err := userRepo.GetAllUsers(ctx)
				if err != nil {
					log.Println("Error collecting users:", err)
					continue
				}

				for _, user := range users {
					assignBadges(ctx, user.ID, postRepo, promiseRepo, followerRepo, badgeService)
				}
				log.Printf("Watcher finished cycle")
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func assignBadges(
	ctx context.Context,
	userID uuid.UUID,
	postRepo post.PostRepository,
	promiseRepo promise.PromiseRepository,
	followerRepo follower.FollowerRepository,
	badgeService services.BadgeService,
) {
	postCount, _ := postRepo.CountUserPosts(ctx, userID)
	promiseCount, _ := promiseRepo.CountAllPromisesByUserID(ctx, userID)
	friendCount, _ := followerRepo.CountMutualFriends(ctx, userID)

	badgeConditions := []struct {
		Code  string
		Check bool
	}{
		{"post_1", postCount >= 1},
		{"post_10", postCount >= 10},
		{"post_50", postCount >= 50},
		{"promise_1", promiseCount >= 1},
		{"promise_5", promiseCount >= 5},
		{"promise_15", promiseCount >= 15},
		{"friends_1", friendCount >= 1},
		{"friends_5", friendCount >= 5},
		{"friends_15", friendCount >= 15},
	}

	for _, badge := range badgeConditions {
		if !badge.Check {
			continue
		}

		has, err := badgeService.HasUserBadgeByCode(ctx, userID, badge.Code)
		if err != nil {
			log.Printf("Error while assigning badge %s for %s: %v", badge.Code, userID, err)
			continue
		}
		if has {
			continue
		}

		err = badgeService.AssignBadgeByCode(ctx, userID, badge.Code)
		if err != nil {
			log.Printf("Error while assigning badge %s for %s: %v", badge.Code, userID, err)
		} else {
			log.Printf("Badge %s assigned to user %s", badge.Code, userID)
		}
	}
}
