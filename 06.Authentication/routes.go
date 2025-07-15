package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {
	api := app.Group("/api")

	userRouter := api.Group("/auth")
	userRouter.POST("/login", HandleLogin)
	userRouter.POST("/register", HandleRegister)
	userRouter.POST("/logout", HandleLogout)
	userRouter.GET("/me", AuthRequired, FetchMyProfile)
}
