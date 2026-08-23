package authorization

import (
	"jwt/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(d SignUpRequestDTO) error {
	user := models.User{Name: d.Name, Email: d.Email, Password: d.Password}
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthRepository) GetUser(d LoginInRequestDTO) (*models.User, error) {
	var user models.User
	db := r.db

	if d.Email != "" {
		db = db.Where("email = ?", d.Email)
	}

	if d.Name != "" {
		db = db.Where("name = ?", d.Name)
	}

	if err := db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
