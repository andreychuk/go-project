package usecase

import (
	"errors"
	"project/internal/domain/models"
	"project/internal/infrastructure/repository"
)

type RoleUsecase interface {
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(uid uint) error
	GetRoleByID(uid uint) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
}

type roleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewRoleUsecase(rRepo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepo: rRepo,
	}
}

func (r *roleUsecase) CreateRole(role *models.Role) error {
	return r.roleRepo.Create(role)
}

func (r *roleUsecase) UpdateRole(role *models.Role) error {
	return r.roleRepo.Update(role)
}

func (r *roleUsecase) DeleteRole(uid uint) error {
	return r.roleRepo.Delete(uid)
}

func (r *roleUsecase) GetRoleByID(uid uint) (*models.Role, error) {
	role, err := r.roleRepo.FindByID(uid)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role not found")
	}
	return role, nil
}

func (r *roleUsecase) GetAllRoles() ([]models.Role, error) {
	return r.roleRepo.FindAll()
}
