package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	Tag        Tag    `json:"tag"`
	TagId      int    `json:"tag_id"`
	Title      string `json:"title"`
	State      int    `json:"state"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifieOn", time.Now().Unix())
	return nil
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (article []Article) {
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id=", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func ExistArticleByName(name string) bool {
	var article Article
	db.Select("id").Where("name = ?", name).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:      data["tag_id"].(int),
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		CreatedBy:  data["create_by"].(string),
		ModifiedBy: data["modified_by"].(string),
		State:      data["state"].(int),
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}
