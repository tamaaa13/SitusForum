package memberships

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/fastcampus/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.SignUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.SignUpResponse{
		Email:     request.Email,
		Username:  request.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBY: request.Email,
		UpdatedBy: request.Email,
	}

	c.JSON(http.StatusOK, response)
}
