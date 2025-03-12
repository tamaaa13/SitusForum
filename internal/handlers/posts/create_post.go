package posts

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/fastcampus/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	err := h.postSvc.CreatePost(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := posts.PostModel{
		UserID:       userID,
		PostTitle:    request.PostTitle,
		PostContent:  request.PostContent,
		PostHashtags: strings.Join(request.PostHashtags, ","),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBY:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	c.JSON(http.StatusOK, response)
}
