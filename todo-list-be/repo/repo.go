package repo

import (
	"math"
	"todo-list-be/dto"

	"gorm.io/gorm"
)

type Repo[T any] struct {
	DB *gorm.DB
}

func (r *Repo[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repo[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repo[T]) Updates(db *gorm.DB, entity *T, fields any) error {
	return db.Model(entity).Updates(fields).Error
}

func (r *Repo[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repo[T]) CountById(db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *Repo[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}

func (r *Repo[T]) populatePageResult(tx *gorm.DB, entities *[]T, page *dto.PageMetadata, filters ...func(*gorm.DB)*gorm.DB) error{
	var total int64
	err := tx.Model(entities).Scopes(filters...).Count(&total).Error
	if err != nil {
		return err
	}
	page.TotalItem = uint(total)

	page.TotalPage = uint(math.Ceil(float64(total) / float64(page.Size)))
	return nil
}