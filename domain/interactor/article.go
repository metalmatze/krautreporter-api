package interactor

import "github.com/MetalMatze/Krautreporter-API/domain/entity"

type ArticleInteractor struct {
	ArticleRepository entity.ArticleRepository
}

func (i ArticleInteractor) SaveAll(authors []entity.Article) error {
	return i.ArticleRepository.SaveAll(authors)
}

func (i ArticleInteractor) FindByID(id int) (*entity.Article, error) {
	return i.ArticleRepository.FindByID(id)
}
