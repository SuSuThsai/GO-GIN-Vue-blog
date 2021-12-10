package model

import (
	"GO-GIN-Vue-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"math"
)

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `gorm:"type:varchar(500);notnull;"json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2"json:"status"`
}

// AddComment 新增评论
func AddComment(data *Comment) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetComment 查询单个评论
func GetComment(id int) (Comment, int) {
	var comment Comment
	err = db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCESS
}

// GetCommentList 获取后台所有评论列表
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&commentList).Count(&total)
	err = db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

// GetCommentCount 获取文章评论数量
func GetCommentCount(id int) (int, int64) {
	var comment Comment
	var total int64
	err = db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	if err != nil {
		return errmsg.ERROR, math.MaxInt64
	}
	return errmsg.SUCCESS, total
}

// GetCommentListFront 展示评论列表
func GetCommentListFront(id int, pageSIze, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSIze).Offset((pageNum-1)*pageSIze).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

// DeleteComment 删除评论
func DeleteComment(id uint) int {
	var comment Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//修改评论，感觉不需要后面要再做

// CheckComment 通过评论
func CheckComment(id int, data *Comment) int {
	var comment Comment
	var ans Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Update(maps).First(&ans).Error
	db.Model(&article).Where("id = ?", ans.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count+?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UncheckComment 驳回评论
func UncheckComment(id int, data *Comment) int {
	var comment Comment
	var ans Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Update(maps).First(&ans).Error
	db.Model(&article).Where("id = ?", ans.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
