package app

import (
	"github.com/gin-gonic/gin"
)

// Handler type is routes handler
type Handler struct {
	Config Config
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.NoRoute(h.notFound)
	router.NoMethod(h.notFound)
	router.LoadHTMLGlob("app/web/templates/*.html")
	router.Static("/static", "app/web/static")

	router.Use(gin.Logger())

	router.GET("/", h.redirectToHome) // Redirects all requests to a /home path

	home := router.Group("/home")
	{
		home.GET("/", h.home)
	}

	qs := router.Group("/qs")
	{
		qs.GET("/", h.redirectToProfile)
		qs.GET("/todayQuestion", h.getTodayQuestion)
		qs.GET("/question", h.question)

		profile := qs.Group("/profile")
		{
			profile.GET("/", h.redirectToProfile)
			profile.GET("/:name", h.profile)
			profile.GET("/verifyEmail", h.verifyEmail)
			profile.GET("/login", h.login)
			profile.GET("/logout", h.logout)
		}
	}

	api := router.Group("/api/" + h.Config.Api.Version)
	{
		api.GET("/", h.getApiInfo)
		api.POST("/", h.updateServer)
		api.GET("/docs", h.apiDocs)
		api.GET("/getQuestion", h.getQuestion)
		//api.POST("/getQuestion", h.getQuestion)
	}

	return router
}
