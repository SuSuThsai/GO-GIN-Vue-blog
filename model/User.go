package model

import (
	"GO-GIN-Vue-blog/utils/errmsg"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null"json:"password"validate:"required,min=6,max=20" label:"密码""`
	Role     int    `gorm:"type:int;DEFAULT:2"json:"role"validate:"required,gte=2" label:"角色"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	} else if users.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST //10003
	}
	return errmsg.SUCCESS
}

// CreatUser 创建用户
func CreatUser(data *User) (code int) {
	//data.Password=ScryptPW(data.Password)
	err = db.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	fmt.Println(id)
	err = db.Where("id = ?", id).Delete(&user).Error
	//对不存在id没有进行处理
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeSave ScryptPW 密码加密 密码加密方法有：bcrypt，scrypt，加salt hash
func (u *User) BeforeSave() {
	u.Password = ScryptPW(u.Password)
}
func ScryptPW(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{255, 22, 18, 33, 99, 66, 25, 11}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

// ValidateLogin 登陆验证
func ValidateLogin(username string, password string) (User, int) {
	var user User
	//var PasswordErr error

	db.Where("username = ?", username).First(&user)

	//PasswordErr=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	//if PasswordErr != nil {
	//	return user,errmsg.ERROR_PASSWORD_WRONG
	//}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NOT_RIGHT
	}
	return user, errmsg.SUCCESS
}
