package controller

import (
	"net/http"
	"strconv"

	"github.com/MetalMatze/Krautreporter-API/http/marshaller"
	"github.com/MetalMatze/Krautreporter-API/krautreporter/repository"
	"github.com/gin-gonic/gin"
	"github.com/gollection/gollection/router"
)

type ArticlesController struct {
	*Controller
}

func (c *ArticlesController) GetArticles(res router.Response, req router.Request) error {
	id := repository.MaxArticleID
	if req.Query("olderthan") != "" {
		olderthan, err := strconv.Atoi(req.Query("olderthan"))
		if err != nil {
			c.log.Info("Can't convert olderthan id to int", "err", err.Error())
			return res.AbortWithStatus(http.StatusInternalServerError)
		}
		id = olderthan
	}

	articles, err := c.interactor.ArticlesOlderThan(id, 20)
	if err != nil {
		if err == repository.ErrArticleNotFound {
			c.log.Debug("Can't find olderthan article", "id", id)
			status := http.StatusNotFound
			return res.JSON(status, gin.H{
				"message":     http.StatusText(status),
				"status_code": status,
			})
		}

		c.log.Warn("Failed to get olderthan articles", "id", id, "err", err)
		return res.AbortWithStatus(http.StatusInternalServerError)
	}

	return res.JSON(http.StatusOK, marshaller.Articles(articles))
}

func (c *ArticlesController) GetArticle(res router.Response, req router.Request) error {
	id, err := strconv.Atoi(req.Param("id"))
	if err != nil {
		c.log.Info("Can't convert article id to int", "err", err.Error())
		return res.AbortWithStatus(http.StatusInternalServerError)
	}

	article, err := c.interactor.ArticleByID(id)
	if err != nil {
		c.log.Debug("Can't find article", "id", id)
		return res.AbortWithStatus(http.StatusNotFound)
	}

	return res.JSON(http.StatusOK, marshaller.Article(article))
}
