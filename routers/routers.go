package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/api/v1"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Runmode)
	apiV1 := r.Group("/v1/api")
	{
		apiV1.GET("tags", v1.GetTags)
		apiV1.POST("addTag", v1.AddTag)
		apiV1.PUT("editTag", v1.EditTag)
		apiV1.DELETE("deleteTag", v1.DeleteTag)
	}
	apiV2 := r.Group("/v1/api")
	{
		apiV2.GET("article", v1.GetArticle)
		apiV2.GET("articles", v1.GetArticles)
		apiV2.POST("addArticle", v1.AddArticle)
		apiV2.PUT("editArticle", v1.EditArticle)
		apiV2.DELETE("deleteArticle", v1.DeleteArticle)
	}
	return r
}
