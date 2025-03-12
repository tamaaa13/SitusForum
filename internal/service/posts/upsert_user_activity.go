package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tamaaa13/fastcampus/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBY: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postsRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user from database")
		return err
	}

	if userActivity == nil {
		// create user activity
		if !req.IsLiked {
			return errors.New("you don't like it yet")
		}
		err = s.postsRepo.CreateUserActivity(ctx, model)
	} else {
		// update user activity
		err = s.postsRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or update user activitiy to database")
	}

	return nil

}
