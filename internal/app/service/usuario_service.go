package service

import (
	"crud-golang/internal/app/model"
	"crud-golang/internal/app/repository"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Atribuir a senha criptografada ao usuário
	user.Password = string(hashedPassword)

	user, err = s.userRepo.Create(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return user, nil
}

func (s *UserService) GetUserByID(id int64) (*model.User, error) {

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}

func (s *UserService) Update(user *model.User) (*model.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Atribuir a senha criptografada ao usuário
	user.Password = string(hashedPassword)

	user, err = s.userRepo.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return user, nil
}

func (s *UserService) Delete(id int64) error {

	err := s.userRepo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "failed delete user")
	}

	return nil
}
