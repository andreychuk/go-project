package repository

import (
	"project/internal/domain/models"

	"gorm.io/gorm"
)

type GroupRepository interface {
	Create(group *models.Group) error
	Update(group *models.Group) error
	Delete(uid uint) error
	FindByID(uid uint) (*models.Group, error)
	FindAll() ([]models.Group, error)
	FindByIDs(ids []uint) ([]*models.Group, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *groupRepository) Update(group *models.Group) error {
	return r.db.Save(group).Error
}

func (r *groupRepository) Delete(uid uint) error {
	return r.db.Delete(&models.Group{}, uid).Error
}

func (r *groupRepository) FindByID(uid uint) (*models.Group, error) {
	var group models.Group
	err := r.db.Preload("Users").First(&group, uid).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) FindAll() ([]models.Group, error) {
	var groups []models.Group
	err := r.db.Preload("Users").Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *groupRepository) FindByIDs(ids []uint) ([]*models.Group, error) {
	var groups []*models.Group
	err := r.db.Find(&groups, ids).Error
	return groups, err
}
