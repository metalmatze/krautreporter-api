package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/MetalMatze/Krautreporter-API/krautreporter/entity"
	"github.com/gollection/gollection/cache"
	"github.com/gollection/gollection/log"
	"github.com/jinzhu/gorm"
)

var ErrAuthorNotFound = errors.New("Author not found")

type GormAuthorRepository struct {
	repository
}

func NewGormAuthorRepository(c cache.Cache, db *gorm.DB, log log.Logger) *GormAuthorRepository {
	return &GormAuthorRepository{repository: newRepository(c, db, log)}
}

func (r GormAuthorRepository) Find() ([]*entity.Author, error) {
	if cached, exists := r.cache.Get("authors.list"); exists {
		return cached.([]*entity.Author), nil
	}

	var authors []*entity.Author

	r.db.Preload("Images").Order("ordering desc").Find(&authors)

	r.cache.Set("authors.list", authors, time.Minute)

	return authors, nil
}

func (r GormAuthorRepository) FindByID(id int) (*entity.Author, error) {
	if cached, exists := r.cache.Get(fmt.Sprintf("authors.%d", id)); exists {
		return cached.(*entity.Author), nil
	}

	var author entity.Author
	r.db.Preload("Images").First(&author, "id = ?", id)

	if author.ID == 0 {
		return nil, ErrAuthorNotFound
	}

	r.cache.Set(fmt.Sprintf("authors.%d", author.ID), &author, time.Minute)

	return &author, nil
}

func (r GormAuthorRepository) SaveAll(authors []entity.Author) error {
	tx := r.db.Begin()
	for _, a := range authors {
		author := entity.Author{ID: a.ID}
		tx.Preload("Crawl").Preload("Images").FirstOrCreate(&author)

		author.Ordering = a.Ordering
		author.Name = a.Name
		author.Title = a.Title
		author.URL = a.URL

		for _, i := range a.Images {
			author.AddImage(i)
		}

		if author.Crawl.ID == 0 {
			author.Crawl = entity.Crawl{Next: time.Now()}
		}

		tx.Save(&author)
	}
	tx.Commit()

	return nil
}

func (r GormAuthorRepository) Save(author entity.Author) error {
	if result := r.db.Save(&author); result.Error != nil {
		return result.Error
	}

	return nil
}
