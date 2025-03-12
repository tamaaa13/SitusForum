package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tamaaa13/fastcampus/internal/model/posts"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()

	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBY:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postsRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create comment to repository")
		return err
	}

	return nil
}
