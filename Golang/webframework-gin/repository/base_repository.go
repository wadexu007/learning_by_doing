package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository interface {
	Update(model any, fields ...string) (int64, error)
	First(model any, id string, includeDeleted bool) error
	Delete(model any) (int64, error)
	Create(model any) error
}

type baseRepository struct {
	DB *gorm.DB
}

func (b *baseRepository) Update(model any, fields ...string) (int64, error) {
	var result *gorm.DB
	if len(fields) == 0 {
		result = b.DB.Updates(model)
	} else {
		result = b.DB.Select(fields).Updates(model)
	}
	return result.RowsAffected, result.Error
}
func (b *baseRepository) First(model any, id string, includeDeleted bool) error {
	uuid, error := uuid.Parse(id)
	if error != nil {
		return error
	}
	var result *gorm.DB
	if includeDeleted {
		result = b.DB.Unscoped().First(model, uuid)
	} else {
		result = b.DB.First(model, uuid)
	}
	return result.Error
}

func (b *baseRepository) Delete(model any) (int64, error) {
	updateResult := b.DB.Updates(model)
	if updateResult.Error != nil || updateResult.RowsAffected == 0 {
		return 0, updateResult.Error
	}
	deleteResult := b.DB.Delete(model)
	return deleteResult.RowsAffected, deleteResult.Error
}

func (b *baseRepository) Create(model any) error {
	result := b.DB.Create(model)
	return result.Error
}
