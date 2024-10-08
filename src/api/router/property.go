package router

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/handler"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}
