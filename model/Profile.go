package model

import "GO-GIN-Vue-blog/utils/errmsg"

type Profile struct {
	ID        int    `gorm:"primaryKey"json:"id"`
	Name      string `gorm:"type:varchar(20)"json:"name"`
	Desc      string `gorm:"type:varchar(200)"json:"desc"`
	Qqchat    string `gorm:"type:varchar(200)"json:"qq_chat"`
	Wechat    string `gorm:"type:varchar(100)"json:"wechat"`
	Twitter   string `gorm:"type:varchar(200)"json:"twitter"`
	Line      string `gorm:"type:varchar(200)"json:"line"`
	Weibo     string `gorm:"type:varchar(200)"json:"weibo"`
	Bili      string `gorm:"type:varchar(200)"json:"bili"`
	Email     string `gorm:"type:varchar(200)"json:"email"`
	Img       string `gorm:"type:varchar(200)"json:"img"`
	Avatar    string `gorm:"type:varchar(200)"json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)"json:"icp_record"`
}

// GetProfile 获取个人信息设置
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

// UpdateProfile 更新个人设置
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
