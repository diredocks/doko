package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex;not null" json:"name"`
	Books []*Book `gorm:"many2many:book_tags;" json:"books"`
}

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) base() *gorm.DB {
	return r.db.Model(&Tag{}).Preload("Books")
}

func (r *TagRepository) Create(tag *Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) GetAll() ([]*Tag, error) {
	var tags []*Tag
	err := r.db.Model(&Tag{}).Find(&tags).Error
	return tags, err
}

func (r *TagRepository) GetByID(id uint) (*Tag, error) {
	var tag Tag
	err := r.base().First(&tag, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) GetByName(name string) (*Tag, error) {
	var tag Tag
	err := r.base().FirstOrCreate(&tag, Tag{Name: name}).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) Delete(id uint) error {
	return r.db.Delete(&Tag{}, "id = ?", id).Error
}
