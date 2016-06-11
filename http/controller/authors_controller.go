package controller

import (
	"net/http"
	"strconv"

	"github.com/MetalMatze/Krautreporter-API/http/marshaller"
	"github.com/MetalMatze/Krautreporter-API/krautreporter/entity"
	"github.com/MetalMatze/Krautreporter-API/krautreporter/repository"
	"github.com/gin-gonic/gin"
	"github.com/gollection/gollection/router"
)

type AuthorInteractor interface {
	GetAll() ([]*entity.Author, error)
	FindByID(id int) (*entity.Author, error)
}

type AuthorsController struct {
	Controller
	AuthorInteractor AuthorInteractor
}

func (c *AuthorsController) GetAuthors(res router.Response, req router.Request) error {
	authors, err := c.AuthorInteractor.GetAll()
	if err != nil {
		c.log.Info("Can't get all authors", "err", err.Error())
		return res.AbortWithStatus(http.StatusInternalServerError)
	}

	return res.JSON(http.StatusOK, marshaller.Authors(authors))
}

func (c *AuthorsController) GetAuthor(res router.Response, req router.Request) error {
	id, err := strconv.Atoi(req.Param("id"))
	if err != nil {
		c.log.Info("Can't convert author id to int", "err", err.Error())
		return res.AbortWithStatus(http.StatusInternalServerError)
	}

	author, err := c.AuthorInteractor.FindByID(id)
	if err != nil {
		if err == repository.ErrAuthorNotFound {
			c.log.Debug("Can't find author", "id", id)
			status := http.StatusNotFound
			return res.JSON(status, gin.H{"message": http.StatusText(status), "status_code": status})
		}

		c.log.Warn("Can't get author", "id", id, "err", err.Error())
		return res.AbortWithStatus(http.StatusInternalServerError)
	}

	return res.JSON(http.StatusOK, marshaller.Author(author))
}
