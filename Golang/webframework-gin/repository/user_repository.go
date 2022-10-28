package repository

import (
	. "cost-analyzer/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	SearchUser(name string, includeDeleted bool) ([]*User, error)
	BaseRepository
}

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepository := &userRepository{baseRepository{DB: db}}
	return userRepository
}

type userRepository struct {
	baseRepository
}

func (r *userRepository) SearchUser(name string, includeDeleted bool) ([]*User, error) {
	var result []*User
	tx := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&result)
	if includeDeleted {
		tx = tx.Unscoped()
	}
	return result, tx.Error
}
