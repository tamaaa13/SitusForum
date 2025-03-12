package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/fastcampus/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.LoginResponse{
		Email:        request.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
