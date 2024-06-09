package middleware

import (
	"errors"
	"github.com/SajjadRezaei/go-clean-architecture/pkg/limiter"
	"net/http"
	"time"

	"github.com/SajjadRezaei/go-clean-architecture/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		getLimiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !getLimiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
