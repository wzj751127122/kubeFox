package authory

import (
	"k8s-platform/logic"
	"k8s-platform/middle"
	"k8s-platform/model"

	"github.com/gin-gonic/gin"
)

var AuthoritiyController authorityController

type authorityController struct{}

func (a *authorityController) GetAuthorityList(ctx *gin.Context) {
	params := &model.PageInfo{}
	if err := ctx.ShouldBind(params); err != nil {
		
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	data, err := logic.GetAuthorityList(ctx, *params)
	if err != nil {
		
		middle.ResponseError(ctx,middle.CodeInvalidParam)
		return
	}
	middle.ResponseSuccess(ctx, data)

}


// GetPolicyPathByAuthorityId
// @Tags      Casbin
// @Summary   获取权限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.CasbinInReceive                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  middle.ResponseData{msg=string}  "获取权限列表,返回包括casbin详情列表"
// @Router    /api/authority/getPolicyPathByAuthorityId [get]
func (a *authorityController) GetPolicyPathByAuthorityId(ctx *gin.Context) {
	rule := &model.CasbinInReceive{}
	if err := ctx.ShouldBind(rule); err != nil {
		
		middle.ResponseError(ctx,middle.CodeInvalidParam)
		return
	}
	middle.ResponseSuccess(ctx, logic.GetPolicyPathByAuthorityId(rule.AuthorityId))
}

// UpdateCasbinByAuthorityId
// @Tags      Casbin
// @Summary   通过角色更新接口权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.UpdateCasbinInput                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  middle.ResponseData{msg=string}  "通过角色更新接口权限"
// @Router    /api/authority/updateCasbinByAuthority [post]
func (a *authorityController) UpdateCasbinByAuthorityId(ctx *gin.Context) {
	params := &model.UpdateCasbinInput{}
	if err := ctx.ShouldBind(params); err != nil {
	
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	if err := logic.UpdateCasbin(params.AuthorityId, params.CasbinInfo); err != nil {
		
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "")
}


