package menu

import (
	"fmt"
	"k8s-platform/logic"
	"k8s-platform/middle"
	"k8s-platform/model"
	"k8s-platform/utils"

	"github.com/gin-gonic/gin"
)

var MenuController menuController
type menuController struct{}

// GetMenusByAuthID
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      dto.Empty                                                  true  "空"
// @Success   200   {object}  middle.Response{data=dto.SysBaseMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /api/menu/:authID/getMenuByAuthID [get]
func (m *menuController) GetMenusByAuthID(ctx *gin.Context) {
	authID, err := utils.ParseUint(ctx.Param("authID"))
	if err != nil || authID == 0 {

		middle.ResponseErrorWithMsg(ctx, middle.CodeInvalidParam, fmt.Errorf("authID empty"))
	}
	menus, err := logic.GetMenuByAuthorityID(ctx, authID)
	if err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, &model.SysMenusResponse{Menus: menus})
}

// GetBaseMenus
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      dto.Empty                                                  true  "空"
// @Success   200   {object}  middle.Response{data=dto.SysBaseMenusResponse,msg=string}  "获取系统菜单详情列表"
// @Router    /api/menu/getBaseMenuTree [get]
func (m *menuController) GetBaseMenus(ctx *gin.Context) {
	menus, err := logic.GetBassMenu(ctx)
	if err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, menus)
}

// AddBaseMenu
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  middle.Response{msg=string}  "新增菜单"
// @Router    /api/menu/add_base_menu [post]
func (m *menuController) AddBaseMenu(ctx *gin.Context) {
	params := &model.AddSysMenusInput{}
	if err := ctx.ShouldBind(params); err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
	}
	if err := logic.AddBaseMenu(ctx, params); err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "添加成功")
}

// AddMenuAuthority
// @Tags      AuthorityMenu
// @Summary   增加menu和角色关联关系
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.AddMenuAuthorityInput  true  "角色ID"
// @Success   200   {object}  response.Response{msg=string}   "增加menu和角色关联关系"
// @Router    /api/menu/add_menu_authority [post]
func (m *menuController) AddMenuAuthority(ctx *gin.Context) {
	params := &model.AddMenuAuthorityInput{}
	if err := ctx.ShouldBind(params); err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
	}
	if err := logic.AddMenuAuthority(ctx, params.Menus, params.AuthorityId); err != nil {

		middle.ResponseError(ctx,middle.CodeServerBusy)
	}
	middle.ResponseSuccess(ctx, "添加成功")
}
