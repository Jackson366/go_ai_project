package dao

import (
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
)

func (d *Dao) SelectOne(userAccount, userPassword string) (model.User, error) {
	user := model.User{UserAccount: userAccount, UserPassword: userPassword}
	return user.Get(d.engine)
}

func (d *Dao) GetById(id uint64) (model.User, error) {
	user := model.User{Model: &model.Model{ID: id}}
	return user.GetById(d.engine)
}

func (d *Dao) CountUser(account string) (int, error) {
	user := model.User{UserAccount: account}
	return user.Count(d.engine)
}

func (d *Dao) CreateUser(user *model.User) error {
	return user.Create(d.engine)
}

func (d *Dao) UpdateUser(id uint64, userName, userAvatar, userProfile, userRole string) error {
	user := model.User{Model: &model.Model{ID: id}}
	values := make(map[string]interface{})
	// 判断userName,userAvatar,userProfile,userRole是否为空，不为空则更新
	if userName != "" {
		values["userName"] = userName
	}
	if userAvatar != "" {
		values["userAvatar"] = userAvatar
	}
	if userProfile != "" {
		values["userProfile"] = userProfile
	}
	if userRole != "" {
		values["userRole"] = userRole
	}
	return user.Update(d.engine, values)
}

func (d *Dao) DeleteUser(id uint64) error {
	user := model.User{Model: &model.Model{ID: id}}
	return user.Delete(d.engine)
}

func (d *Dao) GetUserList(id uint64, unionId, mpOpenId, userName, userProfile, userRole string, page, pageSize int) ([]*model.User, error) {
	user := model.User{
		Model:       &model.Model{ID: id},
		UnionId:     unionId,
		MpOpenId:    mpOpenId,
		UserName:    userName,
		UserProfile: userProfile,
		UserRole:    userRole,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountUserList(id uint64, unionId, mpOpenId, userName, userProfile, userRole string) (int, error) {
	user := model.User{
		Model:       &model.Model{ID: id},
		UnionId:     unionId,
		MpOpenId:    mpOpenId,
		UserName:    userName,
		UserProfile: userProfile,
		UserRole:    userRole,
	}
	return user.Count(d.engine)
}
