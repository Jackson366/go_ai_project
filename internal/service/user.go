package service

import (
	"goAiproject/global"
	"goAiproject/internal/model"
	"goAiproject/pkg/app"
	"goAiproject/pkg/errcode"
	"goAiproject/pkg/util"
	"time"
)

type UserLoginRequest struct {
	UserAccount  string `form:"userAccount" binding:"required,min=4,max=100"`
	UserPassword string `form:"userPassword" binding:"required,min=8,max=30"`
}

type UserRegisterRequest struct {
	UserAccount   string `form:"userAccount" binding:"required,min=4,max=100"`
	UserPassword  string `form:"userPassword" binding:"required,min=8,max=30"`
	CheckPassword string `form:"checkPassword" binding:"required,min=8,max=30"`
}

type LoginUserInfo struct {
	ID          uint64    `json:"id"`
	UserName    string    `json:"userName"`
	UserAvatar  string    `json:"userAvatar"`
	UserProfile string    `json:"userProfile"`
	UserRole    string    `json:"userRole"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type UserInfo struct {
	ID          uint64
	UserName    string
	UserAvatar  string
	UserProfile string
	UserRole    string
	CreateTime  time.Time
}

type UserAddRequest struct {
	UserName    string `form:"userName" binding:"required,min=2,max=100"`
	UserAccount string `form:"userAccount" binding:"required,min=4,max=100"`
	UserAvatar  string `form:"userAvatar" binding:"max=1024"`
	UserRole    string `form:"userRole" binding:"required,min=2,max=100"`
}

type UserUpdateRequest struct {
	ID          uint64 `form:"id"`
	UserName    string `form:"userName" binding:"min=2,max=100"`
	UserAvatar  string `form:"userAvatar" binding:"max=1024"`
	UserProfile string `form:"userProfile" binding:"max=512"`
	UserRole    string `form:"userRole" binding:"max=100"`
}

type UserUpdateMyRequest struct {
	UserName    string `form:"userName" binding:"min=2,max=100"`
	UserAvatar  string `form:"userAvatar" binding:"max=1024"`
	UserProfile string `form:"userProfile" binding:"max=512"`
}

type UserQueryRequest struct {
	ID          uint64 `form:"id" binding:"max=100"`
	UnionId     string `form:"unionId" binding:"max=100"`
	MpOpenId    string `form:"mpOpenId" binding:"max=100"`
	UserName    string `form:"userName" binding:"max=100"`
	UserProfile string `form:"userProfile" binding:"max=1024"`
	UserRole    string `form:"userRole" binding:"max=100"`
}

func (svc *Service) UserLogin(param *UserLoginRequest) (string, error) {
	encryptPassword := util.EncodeSha256WithSalt(param.UserPassword)
	user, err := svc.dao.SelectOne(param.UserAccount, encryptPassword)
	if err != nil {
		global.Logger.Fatalf("用户不存在或密码错误,账户: %s", param.UserAccount)
		return "", errcode.ErrorUserOrPasswordNotExist
	}
	token, err := app.GenerateToken(user.ID)
	if err != nil {
		global.Logger.Fatalf("生成token失败,账户: %s", param.UserAccount)
		return "", errcode.UnauthorizedTokenGenerate
	}
	return token, nil
}

func (svc *Service) UserRegister(param *UserRegisterRequest) (string, error) {
	password := param.UserPassword
	checkPassword := param.CheckPassword
	// 密码和校验密码相同
	if password != checkPassword {
		// 错误抛不出去
		return "", errcode.ErrorPasswordInconsistency
	}
	var count int
	var err error
	count, err = svc.dao.CountUser(param.UserAccount)
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errcode.ErrorUserAccountExist
	}
	encryptPassword := util.EncodeSha256WithSalt(param.UserPassword)
	user := model.User{
		UserAccount:  param.UserAccount,
		UserPassword: encryptPassword,
	}
	err = svc.dao.CreateUser(&user)
	if err != nil {
		return "", err
	}
	return user.UserAccount, nil
}

func (svc *Service) IsAdmin() (bool, error) {
	return false, nil
}

func (svc *Service) JudgeAdmin(user *model.User) (bool, error) {
	return false, nil
}

func (svc *Service) GetUserById(userId uint64) (*model.User, error) {
	user, err := svc.dao.GetById(userId)
	if err != nil {
		return nil, errcode.ErrorUserNotExist
	}
	return &user, nil
}

func (svc *Service) GetUserInfoById(userId uint64) (*UserInfo, error) {
	user, err := svc.dao.GetById(userId)
	if err != nil {
		return nil, errcode.ErrorUserNotExist
	}
	var userInfo = UserInfo{
		ID:          user.ID,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserProfile: user.UserProfile,
		UserRole:    user.UserRole,
		CreateTime:  user.CreateTime,
	}
	return &userInfo, nil
}

func (svc *Service) GetLoginUserInfo(userId uint64) (*LoginUserInfo, error) {
	//var userInfo LoginUserInfo
	user, err := svc.dao.GetById(userId)
	if err != nil {
		return nil, errcode.ErrorUserNotExist
	}
	var userInfo = LoginUserInfo{
		ID:          user.ID,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserProfile: user.UserProfile,
		UserRole:    user.UserRole,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
	return &userInfo, nil
}

func (svc *Service) GetLoginUserPermitNull() (*model.User, error) {
	return nil, nil
}

func (svc *Service) UserLogout() (bool, error) {
	// token失效即可
	return false, nil
}

func (svc *Service) GetUserInfo(users []*model.User) ([]*UserInfo, error) {
	var userInfos []*UserInfo
	for _, user := range users {
		var userInfo = UserInfo{
			ID:          user.ID,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserProfile: user.UserProfile,
			UserRole:    user.UserRole,
			CreateTime:  user.CreateTime,
		}
		userInfos = append(userInfos, &userInfo)
	}
	return userInfos, nil
}

func (svc *Service) AddUser(param *UserAddRequest) error {
	var user = model.User{
		UserName:    param.UserName,
		UserAccount: param.UserAccount,
		UserAvatar:  param.UserAvatar,
		UserRole:    param.UserRole,
	}
	defaultPassword := "12345678"
	user.UserPassword = util.EncodeSha256WithSalt(defaultPassword)
	err := svc.dao.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateUser(param *UserUpdateRequest) error {
	err := svc.dao.UpdateUser(param.ID, param.UserName, param.UserAvatar, param.UserProfile, param.UserRole)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteUser(id uint64) error {
	err := svc.dao.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetUserList(param *UserQueryRequest, pager *app.Pager) ([]*model.User, error) {
	return svc.dao.GetUserList(param.ID, param.UnionId, param.MpOpenId, param.UserName, param.UserProfile, param.UserRole, pager.Page, pager.PageSize)
}

func (svc *Service) CountUserList(param *UserQueryRequest) (int, error) {
	return svc.dao.CountUserList(param.ID, param.UnionId, param.MpOpenId, param.UserName, param.UserProfile, param.UserRole)
}
