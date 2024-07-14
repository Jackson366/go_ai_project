package errcode

var (
	ErrorPasswordInconsistency  = NewError(20010001, "两次密码不一致")
	ErrorUserAccountExist       = NewError(20010002, "用户已存在")
	ErrorUserOrPasswordNotExist = NewError(20010003, "用户不存在或密码错误")
	ErrorUserNotLogin           = NewError(20010004, "用户未登录")
	ErrorUserNotExist           = NewError(20010005, "用户不存在")
	ErrorAddUserFail            = NewError(20010006, "添加用户失败")
	ErrorUpdateUserFail         = NewError(20010007, "更新用户失败")
	ErrorDeleteUserFail         = NewError(20010008, "删除用户失败")
	ErrorCountUserListFail      = NewError(20010009, "统计用户列表失败")
	ErrorGetUserListFail        = NewError(20010010, "获取用户列表失败")
	ErrorUserLogoutFail         = NewError(20010011, "用户登出失败")
)
