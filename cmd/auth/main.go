package main

import (
	"jwt/dependencies"
	"jwt/initializers"
	authorization "jwt/internal/app/auth"
	"jwt/internal/app/hasher"

	"github.com/gin-gonic/gin"
)

func main() {
	{
		initializers.LoadEnv()
		initializers.ConnToDB()
		initializers.MigrateTables()
	}

	HasherRepo := hasher.NewBcryptHasher(10)
	AuthRepository := authorization.NewAuthRepository(initializers.DB)
	AuthService := authorization.NewAuthService(AuthRepository, HasherRepo)

	depenedencies := dependencies.NewDependencies(AuthService)
	controller := setupController(depenedencies)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.Run()
}
