package controller

import (
	authorization "jwt/internal/app/auth"
	"jwt/pkg/jwt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	AuthService *authorization.AuthService
}

func NewController(
	authService *authorization.AuthService,
) *Controller {
	return &Controller{
		AuthService: authService,
	}
}

func (s *Controller) SignUp(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	dto := authorization.SignUpRequestDTO{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	err := s.AuthService.SignUp(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to signup user",
			"message": err.Error(),
		})
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully created user",
		})
	}
}

func (s *Controller) Login(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	dto := authorization.LoginInRequestDTO{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	cred, err := s.AuthService.LoginIn(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to login user",
			"message": err.Error(),
		})
	}

	// generate JWT token
	method, err := jwt.NewJWTMethod(os.Getenv("SIGNING_METHOD"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to create new JWT method",
			"message": err.Error(),
		})
	}

	options := jwt.NewBasicOptions(cred.ID)

	tokenString, err := jwt.GenerateTokenWithOptions(method, options)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to generate token with options",
			"message": err.Error(),
		})
	}

	// set cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully logged in user",
		})
	}
}

// func ChangePassword(c *gin.Context) {
// 	var body struct {
// 		Password string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "failed to read the body",
// 		})
// 		return
// 	}

// 	if body.Password == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "password is required",
// 		})
// 		return
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "failed to hash password",
// 		})
// 		return
// 	}

// 	var oldPass string

// 	initializers.DB.Take(&models.User{}).Where("id = ?", 8).Select("password").Scan(&oldPass)

// 	if err := bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(body.Password)); err == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "cannot use the same password as the old one",
// 		})
// 		return
// 	}

// 	result := initializers.DB.Model(&models.User{}).Where("id = ?", 8).Update("password", string(hash))

// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "failed to update password",
// 			"message": result.Error,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "password updated successfully",
// 		"result":  result,
// 	})
// }

// func LoginIn(c *gin.Context) {
// 	log.Info("jwt.LoginIn request received")
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "failed to read the body",
// 		})
// 		return
// 	}

// 	result, err :=

// }
