package repository

import (
	"project/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(uid uint) error
	FindByID(uid uint) (*models.User, error)
	FindAll() ([]models.User, error)
	AddGroups(user *models.User, groups []*models.Group) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(uid uint) error {
	return r.db.Delete(&models.User{}, uid).Error
}

func (r *userRepository) FindByID(uid uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Groups").Preload("Role").First(&user, uid).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Groups").Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) AddGroups(user *models.User, groups []*models.Group) error {
	return r.db.Model(user).Association("Groups").Append(groups)
}
