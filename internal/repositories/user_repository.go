package repositories

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID uuid.UUID) error
	IsEmailExists(email string) bool
	IsUsernameExists(username string) bool
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) DeleteUser(userID uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", userID).Error
}

func (r *userRepo) IsEmailExists(email string) bool {
	var count int64
	r.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepo) IsUsernameExists(username string) bool {
	var count int64
	r.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}
