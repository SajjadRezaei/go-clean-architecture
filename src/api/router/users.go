package router

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/handler"
	"github.com/SajjadRezaei/go-clean-architecture/api/middleware"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middleware.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
