package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goAiproject/global"
	"goAiproject/internal/service"
	"goAiproject/pkg/app"
	"goAiproject/pkg/convert"
	"goAiproject/pkg/errcode"
	"goAiproject/pkg/redis"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @Summary 用户登录
// @Produce json
// @Param UserAccount body string true "用户账号"
// @Param UserPassword body string true "用户密码"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /login [post]
func (u User) Login(c *gin.Context) {
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	token, err := svc.UserLogin(&param)
	if err != nil {
		global.Logger.Fatalf("svc.UserLogin err: %v", err)
		switch {
		case errors.Is(err, errcode.ErrorUserOrPasswordNotExist):
			response.ToErrorResponse(errcode.ErrorUserOrPasswordNotExist)
		case errors.Is(err, errcode.UnauthorizedTokenGenerate):
			response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		}
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "登录成功", "data": gin.H{"token": token}})
	return
}

// @Summary 用户注册
// @Produce json
// @Param UserAccount body string true "用户账号"
// @Param UserPassword body string true "用户密码"
// @Param ConfirmPassword body string true "确认密码"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"、
// @Router /register [post]
func (u User) Register(c *gin.Context) {
	param := service.UserRegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	userAccount, err := svc.UserRegister(&param)
	if err != nil {
		global.Logger.Fatalf("svc.UserRegister err: %v", err)
		switch {
		case errors.Is(err, errcode.ErrorUserAccountExist):
			response.ToErrorResponse(errcode.ErrorUserAccountExist)
		case errors.Is(err, errcode.ErrorPasswordInconsistency):
			response.ToErrorResponse(errcode.ErrorPasswordInconsistency)
		}
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "注册成功", "data": gin.H{"userAccount": userAccount}})
}

// @Summary 用户登出
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /logout [get]
func (u User) Logout(c *gin.Context) {
	currentUserId, exists := c.Get("userId")
	fmt.Printf("currentUserId: %v\n", currentUserId)
	response := app.NewResponse(c)
	if !exists {
		// 未写入日志
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	userIdStr := fmt.Sprintf("%v", currentUserId)
	token := c.GetHeader("Authorization")
	expireTime, _ := app.GetTokenRemainingTime(token)

	client := redis.NewRedisClient()
	err := client.Set(context.Background(), userIdStr, token, expireTime).Err()
	if err != nil {
		global.Logger.Fatalf("redis.client.Set err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserLogoutFail)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "退出成功"})
}

// @Summary 获取当前用户信息
// @Produce json
// @Success 200 {object} service.LoginUserInfo "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/get/login [get]
func (u User) GetLoginUser(c *gin.Context) {
	currentUserId, exists := c.Get("userId")
	fmt.Printf("currentUserId: %v\n", currentUserId)
	response := app.NewResponse(c)
	if !exists {
		// 未写入日志
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	userId := currentUserId.(uint64)
	svc := service.New(c.Request.Context())
	user, err := svc.GetLoginUserInfo(userId)
	if err != nil {
		global.Logger.Fatalf("svc.GetLoginUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "查询成功", "data": user})
}

// @Summary 更新当前用户信息
// @Produce json
// @Param UserAccount body string true "用户账号"
// @Param UserName body string true "用户名称"
// @Param UserEmail body string true "用户邮箱"
// @Param UserPhone body string true "用户电话"
// @Param UserSex body string true "用户性别"
// @Param UserAge body string true "用户年龄"
// @Param UserAddress body string true "用户地址"
// @Param UserAvatar body string true "用户头像"
// @Param UserDesc body string true "用户描述"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /user/update/my [put]
func (u User) UpdateMyUser(c *gin.Context) {
	param := service.UserUpdateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	currentUserId, _ := c.Get("userId")
	userId := currentUserId.(uint64)
	param.ID = userId
	svc := service.New(c.Request.Context())
	err := svc.UpdateUser(&param)
	if err != nil {
		global.Logger.Fatalf("svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserFail)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "更新成功"})
}

// @Summary 添加用户
// @Produce json
// @Param UserAccount body string true "用户账号"
// @Param UserName body string true "用户名称"
// @Param UserPassword body string true "用户密码"
// @Param UserEmail body string true "用户邮箱"
// @Param UserPhone body string true "用户电话"
// @Param UserSex body string true "用户性别"
// @Param UserAge body string true "用户年龄"
// @Param UserAddress body string true "用户地址"
// @Param UserAvatar body string true "用户头像"
// @Param UserDesc body string true "用户描述"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /user/add [post]
func (u User) AddUser(c *gin.Context) {
	param := service.UserAddRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.AddUser(&param)
	if err != nil {
		global.Logger.Fatalf("svc.AddUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorAddUserFail)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "添加成功"})
}

// @Summary 删除用户
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /user/delete/{id} [delete]
func (u User) DeleteUser(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		// 未写入日志
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	userId := convert.StrTo(param).MustUint64()
	svc := service.New(c.Request.Context())
	err := svc.DeleteUser(userId)
	if err != nil {
		global.Logger.Fatalf("svc.DeleteUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteUserFail)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "删除成功"})
}

// @Summary 更新用户
// @Produce json
// @Param UserAccount body string true "用户账号"
// @Param UserName body string true "用户名称"
// @Param UserPassword body string true "用户密码"
// @Param UserEmail body string true "用户邮箱"
// @Param UserPhone body string true "用户电话"
// @Param UserSex body string true "用户性别"
// @Param UserAge body string true "用户年龄"
// @Param UserAddress body string true "用户地址"
// @Param UserAvatar body string true "用户头像"
// @Param UserDesc body string true "用户描述"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /user/update [put]
func (u User) UpdateUser(c *gin.Context) {
	param := service.UserUpdateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateUser(&param)
	if err != nil {
		global.Logger.Fatalf("svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserFail)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "更新成功"})
}

// @Summary 获取用户
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
func (u User) GetUserById(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		// 未写入日志
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	userId := convert.StrTo(param).MustUint64()
	svc := service.New(c.Request.Context())
	user, err := svc.GetUserById(userId)
	if err != nil {
		global.Logger.Fatalf("svc.GetLoginUser err: %v", err)
		// 没有做出来自业务层的代码错误判断
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "查询成功", "data": user})
}

// @Summary 获取用户信息
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} service.UserInfo "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/get/info/{id} [get]
func (u User) GetUserInfoById(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		// 未写入日志
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	userId := convert.StrTo(param).MustUint64()
	svc := service.New(c.Request.Context())
	userInfo, err := svc.GetUserInfoById(userId)
	if err != nil {
		global.Logger.Fatalf("svc.GetUserInfoById err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	response.ToResponse(gin.H{"code": 200, "msg": "查询成功", "data": userInfo})
}

// @Summary 获取用户列表
// @Produce json
// @Param Page query int true "页码"
// @Param PageSize query int true "每页显示条数"
// @Param UserAccount query string false "用户账号"
// @Param UserName query string false "用户名称"
// @Param UserEmail query string false "用户邮箱"
// @Param UserPhone query string false "用户电话"
// @Param UserSex query string false "用户性别"
// @Param UserAge query string false "用户年龄"
// @Param UserAddress query string false "用户地址"
// @Param UserAvatar query string false "用户头像"
// @Param UserDesc query string false "用户描述"
// @Success 200 {object} model.UserSwagger "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/list [get]
func (u User) ListUserByPage(c *gin.Context) {
	param := service.UserQueryRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUserList(&param)
	if err != nil {
		global.Logger.Fatalf("svc.CountUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountUserListFail)
		return
	}
	users, err := svc.GetUserList(&param, &pager)
	if err != nil {
		global.Logger.Fatalf("svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}
	response.ToResponseList(users, totalRows)
}

// @Summary 获取用户信息列表
// @Produce json
// @Param Page query int true "页码"
// @Param PageSize query int true "每页显示条数"
// @Param UserAccount query string false "用户账号"
// @Param UserName query string false "用户名称"
// @Param UserEmail query string false "用户邮箱"
// @Param UserPhone query string false "用户电话"
// @Param UserSex query string false "用户性别"
// @Param UserAge query string false "用户年龄"
// @Param UserAddress query string false "用户地址"
// @Param UserAvatar query string false "用户头像"
// @Param UserDesc query string false "用户描述"
// @Success 200 {object} model.UserSwagger "成功"
// @Failure 400 {object} errcode.Error "失败"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/list/info [get]
func (u User) ListUserInfoByPage(c *gin.Context) {
	param := service.UserQueryRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUserList(&param)
	if err != nil {
		global.Logger.Fatalf("svc.CountUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountUserListFail)
		return
	}
	users, err := svc.GetUserList(&param, &pager)
	if err != nil {
		global.Logger.Fatalf("svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}
	userInfos, _ := svc.GetUserInfo(users)
	response.ToResponseList(userInfos, totalRows)
}
