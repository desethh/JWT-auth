package main

import (
	"jwt/dependencies"
	"jwt/initializers"
	authorization "jwt/internal/app/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	{
		initializers.LoadEnv()
		initializers.ConnToDB()
		initializers.MigrateTables()
	}

	AuthRepository := authorization.NewAuthRepository(initializers.DB)
	AuthService := authorization.NewAuthService(AuthRepository)

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
	// router.POST("/change-password", controller.ChangePassword)
	router.Run()
}

// func LoginMiddleware() gin.HandlerFunc {
// 	rate := limiter.Rate{
// 		Period: time.Minute,
// 		Limit:  5, // максимум 5 попыток login за минуту с одного IP
// 	}

// 	store := memory.NewStore()
// 	instance := limiter.New(store, rate)

// 	return mgin.NewMiddleware(instance)
// }
