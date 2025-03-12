package posts

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/fastcampus/internal/model/posts"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("postID invalid").Error(),
		})
		return
	}
	userID := c.GetInt64("userID")

	errs := h.postSvc.CreateComment(ctx, postID, userID, request)
	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errs.Error(),
		})
		return
	}

	response := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBY:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	c.JSON(http.StatusOK, response)
}
