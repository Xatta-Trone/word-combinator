package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xatta-trone/words-combinator/controllers/publicController"
	"github.com/xatta-trone/words-combinator/database"
	"github.com/xatta-trone/words-combinator/middlewares"
	"github.com/xatta-trone/words-combinator/repository"
	"github.com/xatta-trone/words-combinator/services"
)

func PublicRoutes(r *gin.Engine) *gin.Engine {
	// init services
	userRepo := repository.NewUserRepository(database.Gdb)
	authService := services.NewAuthService()
	authController := publicController.NewAuthController(userRepo, authService)

	// list controller 

	public := r.Group("/")

	public.POST("/register", authController.Register)
	public.POST("login", authController.Login)

	// public.GET("@:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	ctx.JSON(200, gin.H{"name": name})
	// })

	authRoutes := r.Group("/").Use(middlewares.PublicAuthMiddleware())

	authRoutes.GET("/me", authController.Me)
	authRoutes.PUT("/update", authController.Update)
	authRoutes.PATCH("/update", authController.Update)



	return r
}