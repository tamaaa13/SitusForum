package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tamaaa13/fastcampus/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetails, err := s.postsRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postsRepo.CountLikeByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comments, err := s.postsRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetails: posts.Post{
			ID:           postDetails.ID,
			UserID:       postDetails.UserID,
			Username:     postDetails.Username,
			PostTitle:    postDetails.PostTitle,
			PostContent:  postDetails.PostContent,
			PostHashtags: postDetails.PostHashtags,
			IsLiked:      postDetails.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
