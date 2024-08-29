package usecase

import (
	"errors"
	"project/internal/domain/models"
	"project/internal/infrastructure/repository"
)

type UserUsecase interface {
	CreateUser(user *models.User, groupIDs []uint) error
	UpdateUser(user *models.User, groupIDs []uint) error
	DeleteUser(uid uint) error
	GetUserByID(uid uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userUsecase struct {
	userRepo  repository.UserRepository
	groupRepo repository.GroupRepository
}

func NewUserUsecase(uRepo repository.UserRepository, gRepo repository.GroupRepository) UserUsecase {
	return &userUsecase{
		userRepo:  uRepo,
		groupRepo: gRepo,
	}
}

func (u *userUsecase) CreateUser(user *models.User, groupIDs []uint) error {
	if err := u.userRepo.Create(user); err != nil {
		return err
	}

	if len(groupIDs) > 0 {
		groups, err := u.groupRepo.FindByIDs(groupIDs)
		if err != nil {
			return err
		}
		if err := u.userRepo.AddGroups(user, groups); err != nil {
			return err
		}
	}

	return nil
}

func (u *userUsecase) UpdateUser(user *models.User, groupIDs []uint) error {
	if err := u.userRepo.Update(user); err != nil {
		return err
	}

	if len(groupIDs) > 0 {
		groups, err := u.groupRepo.FindByIDs(groupIDs)
		if err != nil {
			return err
		}
		if err := u.userRepo.AddGroups(user, groups); err != nil {
			return err
		}
	}

	return nil
}

func (u *userUsecase) DeleteUser(uid uint) error {
	return u.userRepo.Delete(uid)
}

func (u *userUsecase) GetUserByID(uid uint) (*models.User, error) {
	user, err := u.userRepo.FindByID(uid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.userRepo.FindAll()
}
