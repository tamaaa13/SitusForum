package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/fastcampus/internal/middleware"
	"github.com/tamaaa13/fastcampus/internal/model/memberships"
)

type membershipSvc interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipSvc
}


func NewHandler(api *gin.Engine, membershipSvc membershipSvc) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}
