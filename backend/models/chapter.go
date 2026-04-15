package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Order   int    `gorm:"not null" json:"order"`
	BookID  uint   `gorm:"not null" json:"book_id"`
}

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}

func (r *ChapterRepository) Create(chapter *Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *ChapterRepository) GetByBookID(bookID uint) ([]*Chapter, error) {
	var chapters []*Chapter
	err := r.db.Where("book_id = ?", bookID).Order("`order` ASC").Find(&chapters).Error
	return chapters, err
}

func (r *ChapterRepository) GetByID(id uint) (*Chapter, error) {
	var chapter Chapter
	err := r.db.First(&chapter, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (r *ChapterRepository) Update(chapter *Chapter) error {
	return r.db.Save(chapter).Error
}

func (r *ChapterRepository) Delete(id uint) error {
	return r.db.Delete(&Chapter{}, "id = ?", id).Error
}
