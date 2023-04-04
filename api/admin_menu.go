package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminMenuList(c *gin.Context) {
	// 绑定并检查参数
	var menuVo request.MenuVo
	var err error
	err = c.ShouldBind(&menuVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	var list []*response.MenuTreeDto
	list, err = menuService.MenuTreeDtoList(menuVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminMenuOption(c *gin.Context) {
	list, err := menuService.MenuTreeOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminMenuUserTree(c *gin.Context) {
	// 从登录状态中获取用户id
	claims := middleware.GetClaims(c)
	if claims == nil {
		middleware.ReportError(c, response.AUTHENTICATION_FAILED, nil)
	}
	list, err := menuService.UserMenuTreeDtoList(claims.ID)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminMenuSave(c *gin.Context) {
	// 绑定并检查参数
	var menuSaveVo request.MenuSaveVo
	var err error
	err = c.ShouldBind(&menuSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 新增数据
	var res bool
	if menuSaveVo.Type == model.MENU_TYPE_GROUP {
		res, err = menuService.SaveMenuGroup(menuSaveVo)
	} else {
		res, err = menuService.SaveMenuItem(menuSaveVo)
	}
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

func AdminMenuRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = menuService.RemoveMenuById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

func AdminMenuUpdate(c *gin.Context) {
	// 绑定并检查参数
	var menuUpdateVo request.MenuUpdateVo
	var err error
	err = c.ShouldBind(&menuUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	menuUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = menuService.UpdateMenu(menuUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}

func AdminMenuUpdateDeleted(c *gin.Context) {
	// 绑定并检查参数
	var deletedVo request.ModelDeletedVo
	var err error
	err = c.ShouldBind(&deletedVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	deletedVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = menuService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	middleware.SetData(c, res)
}