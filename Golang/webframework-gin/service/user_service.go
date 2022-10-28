package service

import (
	. "cost-analyzer/error"
	"cost-analyzer/lib/logger"
	. "cost-analyzer/model"
	. "cost-analyzer/repository"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(id string, includeDeleted bool) (*User, error)
	DeleteUser(id string, deletedBy string) (int64, error)
	UpdateUser(user *User, updatedBy string) (int64, error)
	SearchUser(name string, includeDeleted bool) ([]*User, error)
	CreateUser(user *User) (*uuid.UUID, error)
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{repository: userRepository}
}

type userService struct {
	repository UserRepository
}

func (u *userService) GetUser(id string, includeDeleted bool) (*User, error) {
	user := &User{}
	err := u.repository.First(user, id, includeDeleted)
	//user, err := u.repository.GetUser(id, includeDeleted)
	if err == nil {
		return user, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		logger.ErrorF("failed to get the user because: %s", id, err)
		return nil, InternalServerError
	}
}

func (u *userService) DeleteUser(id string, deletedBy string) (int64, error) {
	//check if user exist before action
	user, err := u.GetUser(id, false)
	if err != nil {
		logger.ErrorF("failed to get user because: %s", id, err)
		return 0, InternalServerError
	} else if user == nil {
		logger.Error("user[%s] doesn't exist", id, err)
		return 0, BadRequestError
	}
	user.UpdatedBy = deletedBy
	//check delete result, and wrap the DB layer error
	rowsAffected, err := u.repository.Delete(&user)
	if err != nil {
		logger.ErrorF("failed to delete the user because: %s", id, err)
		return 0, InternalServerError
	}
	return rowsAffected, nil
}

func (u *userService) UpdateUser(user *User, updatedBy string) (int64, error) {
	//check update result, and wrap the DB layer error
	user.UpdatedBy = updatedBy
	rowsAffected, err := u.repository.Update(user, "name", "email")

	if err != nil {
		logger.ErrorF("failed to update user[%s], because: %s", user.ID.String(), err)
		return 0, InternalServerError
	} else if rowsAffected == 0 {
		logger.ErrorF("user[%s] doesn't exist", user.ID.String())
		return 0, BadRequestError
	} else {
		return rowsAffected, nil
	}

}

func (u *userService) SearchUser(name string, hasDeleted bool) ([]*User, error) {
	users, err := u.repository.SearchUser(name, hasDeleted)
	if err != nil {
		logger.ErrorF("failed to search user because: %s", name, err)
		return nil, InternalServerError
	}

	return users, nil
}

func (u *userService) CreateUser(user *User) (*uuid.UUID, error) {
	err := u.repository.Create(user)
	if err != nil {
		logger.ErrorF("failed to create user because: %s", user, err)
		return nil, InternalServerError
	}
	return &user.ID, nil
}
