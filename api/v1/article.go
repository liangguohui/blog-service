package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/models"
	"github.com/go-programming-tour-book/blog-service/pkg/e"
	"github.com/go-programming-tour-book/blog-service/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state = -1
	code := e.INVALID_PARAMS
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	if !valid.HasErrors() {
		maps["state"] = state
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(c), util.GetPageSize(c), maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {

	}
	log.Printf("id:%s", id)
	valid := validation.Validation{}
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	modifiedBy := c.PostForm("modified_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Required(modifiedBy, "modified_by").Message("不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["content"] = content
			data["desc"] = desc
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:s%,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]string),
	})
}

func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.PostForm("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifieBy := c.PostForm("modifie_by")
	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长100")
	valid.MaxSize(desc, 255, "desc").Message("描述最长255")
	valid.MaxSize(content, 65535, "desc").Message("内容最长65535")
	valid.Required(modifieBy, "modifie_by").Message("修改人不能为空")
	valid.MaxSize(modifieBy, 100, "modifie_by").Message("修改人100个字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				if modifieBy != "" {
					data["modified_by"] = modifieBy
				}
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]string),
	})

}

//删除文章
func DeleteArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "tag_id").Message("标签ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]string),
	})

}
