package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"goAiproject/pkg/app"
)

type User struct {
	*Model
	UserAccount  string `gorm:"type:varchar(256);not null;column:userAccount" json:"userAccount"`
	UserPassword string `gorm:"type:varchar(512);not null;column:userPassword" json:"userPassword"`
	UnionId      string `gorm:"type:varchar(256);column:unionId" json:"unionId"`
	MpOpenId     string `gorm:"type:varchar(256);column:mpOpenId" json:"mpOpenId"`
	UserName     string `gorm:"type:varchar(256);column:userName" json:"userName"`
	UserAvatar   string `gorm:"type:varchar(1024);column:userAvatar" json:"userAvatar"`
	UserProfile  string `gorm:"type:varchar(512);column:userProfile" json:"userProfile"`
	UserRole     string `gorm:"type:varchar(256);default:'user';not null;column:userRole" json:"userRole"`
}

type UserSwagger struct {
	List  []*User
	Pager *app.Pager
}

func (u User) TableName() string {
	return "user"
}

func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("userAccount = ? AND userPassword = ? AND isDelete = ?", u.UserAccount, u.UserPassword, 0)
	err := db.First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}
	return user, nil
}

func (u User) GetById(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("id = ? AND isDelete = ?", u.ID, 0)
	err := db.First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}
	return user, nil
}

func (u User) Count(db *gorm.DB) (int, error) {
	var count int
	if u.UserAccount != "" {
		db = db.Where("userAccount = ?", u.UserAccount)
	}
	if u.UnionId != "" {
		db = db.Where("unionId = ?", u.UnionId)
	}
	if u.MpOpenId != "" {
		db = db.Where("mpOpenId = ?", u.MpOpenId)
	}
	if u.UserName != "" {
		db = db.Where("userName = ?", u.UserName)
	}
	if u.UserProfile != "" {
		db = db.Where("userProfile = ?", u.UserProfile)
	}
	if u.UserRole != "" {
		db = db.Where("userRole = ?", u.UserRole)
	}
	db = db.Where("isDelete = ?", 0)
	err := db.Model(&u).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&u).Where("id = ? AND isDelete = ?", u.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (u User) Delete(db *gorm.DB) error {
	db = db.Where("id = ? AND isDelete = ?", u.ID, 0)
	// 软删除 isDelete = 1
	return db.Delete(&u).Error
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if u.ID > 0 {
		db = db.Where("id = ?", u.ID)
	}
	if u.UnionId != "" {
		db = db.Where("unionId = ?", u.UnionId)
	}
	if u.MpOpenId != "" {
		db = db.Where("mpOpenId = ?", u.MpOpenId)
	}
	if u.UserName != "" {
		db = db.Where("userName = ?", u.UserName)
	}
	if u.UserProfile != "" {
		db = db.Where("userProfile = ?", u.UserProfile)
	}
	if u.UserRole != "" {
		db = db.Where("userRole = ?", u.UserRole)
	}
	db = db.Where("isDelete = ?", 0)
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
