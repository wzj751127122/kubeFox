package user

import (
	"k8s-platform/logic"
	"k8s-platform/middle"
	"k8s-platform/model"
	"k8s-platform/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/login [post]
func (u *userController) Login(ctx *gin.Context) {
	params := new(model.AdminLoginInput)
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	token, err := logic.Login(ctx, params)
	if err != nil {
		middle.ResponseErrorWithMsg(ctx, middle.CodeServerBusy, err)
		return
	}
	middle.ResponseSuccess(ctx, &model.AdminLoginOut{Token: token})
}

// LoginOut godoc
// @Summary 管理员退出登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/loginout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/loginout [get]
func (u *userController) LoginOut(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		logger.Error("服务繁忙")
	}
	cla, _ := claims.(*utils.CustomClaims)
	if err := logic.LoginOut(ctx, cla.ID); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "退出成功")
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  middleware.Response{data=model.SysUser,msg=string}  "获取用户信息"
// @Router    /api/user/getinfo [get]
func (u *userController) GetUserInfo(ctx *gin.Context) {
	clalms, err := utils.GetClaims(ctx)
	if err != nil {
		logger.Error("获取用户信息失败" + err.Error())
		return
	}
	userInfo, err := logic.GetUserInfo(ctx, clalms.ID, clalms.AuthorityId)
	if err != nil {
		logger.Error("获取用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	middle.ResponseSuccess(ctx, userInfo)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.SetUserAuth          true  "角色ID"
// @Success   200   {object}  middleware.Response{msg=string}  "设置用户权限"
// @Router    /api/user/{id}/set_auth [put]
func (u *userController) SetUserAuthority(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	params := new(model.SetUserAuth)
	err = ctx.ShouldBind(params)
	if err != nil {
		logger.Error("ShouldBind失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := logic.SetUserAuth(ctx, uid, params.AuthorityId); err != nil {
		logger.Error("获取用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	// token中存在角色信息，需要生成新的token
	claims := utils.GetUserInfo(ctx)
	claims.AuthorityId = params.AuthorityId
	newToken, err := utils.JwtToken.GenerateToken(claims.BaseClaims)
	if err != nil {
		logger.Error("获取用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	ctx.Header("new-token", newToken)
	ctx.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
	middle.ResponseSuccess(ctx, "操作成功")
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  middleware.Response{msg=string}  "删除用户"
// @Router    /api/user/{id}/delete_user [delete]
func (u *userController) DeleteUser(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("删除用户失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	if err := logic.DeleteUser(ctx, uid); err != nil {
		logger.Error("删除用户失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "操作成功")
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      dto.ChangeUserPwdInput    true  "用户ID, 原密码, 新密码"
// @Success   200   {object}  middleware.Response{msg=string}  "用户修改密码"
// @Router    /api/user/{id}/change_pwd [post]
func (u *userController) ChangePassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("修改用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	params := new(model.ChangeUserPwdInput)
	err = ctx.ShouldBind(params)
	if err != nil {
		logger.Error("ShouldBind失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := logic.ChangePassword(ctx, uid, params); err != nil {
		logger.Error("修改用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "")
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Success   200   {object}  middleware.Response{msg=string}  "重置用户密码"
// @Router    /api/user/{id}/reset_pwd [put]
func (u *userController) ResetPassword(ctx *gin.Context) {
	uid, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		logger.Error("修改用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	if err := logic.ResetPassword(ctx, uid); err != nil {
		logger.Error("修改用户信息失败,参数有误" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "操作成功")
}
