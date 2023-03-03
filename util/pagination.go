package util

//编写分页页码的获取方法
import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * GetPageSize(c)
	} else {
		result = setting.Page
	}
	return result
}

func GetPageSize(c *gin.Context) (result int) {
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	if pageSize > 0 {
		result = pageSize
	} else {
		result = setting.PageSize
	}
	return
}
