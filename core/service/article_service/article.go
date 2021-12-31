package article_service

import (
	"blog-mongo/app/model/request"
	"blog-mongo/app/model/response"
)

type Article interface {
	Get(page int64) []response.Article
	Post(article request.Article) string
	Delete(author string,articleId string) string
	Update(request request.UpdateArticle) string
}
