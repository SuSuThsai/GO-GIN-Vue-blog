package model

import (
	"GO-GIN-Vue-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"json:"title"`
	Cid     int    `gorm:"type:int;not null"json:"cid"`
	Desc    string `gorm:"type:varchar(200)"json:"desc"`
	Content string `gorm:"type:longtext"json:"content"`
	Img     string `gorm:"type:varchar(100)"json:"img"`
}

//// CheckArticle 查询文章是否存在
//func CheckArticle(name string) (code int){
//	var category Category
//	db.Select("id").Where("name = ?",name).First(&category)
//	if category.ID>0 {
//		return errmsg.ERROR_CATEGORYNAME_USED//2001
//	}
//	return errmsg.SUCCESS
//}

// CreatArticle 创建文章
func CreatArticle(article *Article) (code int) {
	//data.Password=ScryptPW(data.Password)
	err = db.Create(article).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// GetCategoryAllArticles 查询分类下的所有文章
func GetCategoryAllArticles(cid int, pageSize int, pageNum int) ([]Article, int, int) {
	var CategoryAllArticles []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&CategoryAllArticles).Count(&total).Error
	if err != nil {
		return CategoryAllArticles, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return CategoryAllArticles, errmsg.SUCCESS, total
}

// GetArticle 查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int) {
	var articles []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	//对不存在id没有进行处理
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
