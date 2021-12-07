package model

import (
	"GO-GIN-Vue-blog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null"json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED //2001
	}
	return errmsg.SUCCESS
}

// CreatCategory 创建分类
func CreatCategory(category *Category) (code int) {
	//data.Password=ScryptPW(data.Password)
	err = db.Create(category).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// GetCategories 查询分类列表
func GetCategories(pageSize int, pageNum int) ([]Category, int) {
	var categories []Category
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categories, total
}

//查询分类下的所有文章

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	fmt.Println(id)
	err = db.Where("id = ?", id).Delete(&category).Error
	//对不存在id没有进行处理
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
