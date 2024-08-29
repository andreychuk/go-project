package usecase

import (
	"errors"
	"project/internal/domain/models"
	"project/internal/infrastructure/repository"
)

type GroupUsecase interface {
	CreateGroup(group *models.Group) error
	UpdateGroup(group *models.Group) error
	DeleteGroup(uid uint) error
	GetGroupByID(uid uint) (*models.Group, error)
	GetAllGroups() ([]models.Group, error)
}

type groupUsecase struct {
	groupRepo repository.GroupRepository
}

func NewGroupUsecase(gRepo repository.GroupRepository) GroupUsecase {
	return &groupUsecase{
		groupRepo: gRepo,
	}
}

func (g *groupUsecase) CreateGroup(group *models.Group) error {
	return g.groupRepo.Create(group)
}

func (g *groupUsecase) UpdateGroup(group *models.Group) error {
	return g.groupRepo.Update(group)
}

func (g *groupUsecase) DeleteGroup(uid uint) error {
	return g.groupRepo.Delete(uid)
}

func (g *groupUsecase) GetGroupByID(uid uint) (*models.Group, error) {
	group, err := g.groupRepo.FindByID(uid)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, errors.New("group not found")
	}
	return group, nil
}

func (g *groupUsecase) GetAllGroups() ([]models.Group, error) {
	return g.groupRepo.FindAll()
}
