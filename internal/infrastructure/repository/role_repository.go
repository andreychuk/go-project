package repository

import (
	"project/internal/domain/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(uid uint) error
	FindByID(uid uint) (*models.Role, error)
	FindAll() ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(uid uint) error {
	return r.db.Delete(&models.Role{}, uid).Error
}

func (r *roleRepository) FindByID(uid uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, uid).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
