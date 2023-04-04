package api

import (
	"blog-server-go/middleware"
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminPermissionList(c *gin.Context) {
	// 绑定并检查参数
	var permissionVo request.PermissionVo
	var err error
	err = c.ShouldBind(&permissionVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	var list []*response.PermissionTreeDto
	list, err = permissionService.PermissionTreeDtoList(permissionVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminPermissionOption(c *gin.Context) {
	list, err := permissionService.PermissionTreeOptionDtoList()
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	middleware.SetData(c, list)
}

func AdminPermissionSave(c *gin.Context) {
	// 绑定并检查参数
	var permissionSaveVo request.PermissionSaveVo
	var err error
	err = c.ShouldBind(&permissionSaveVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 新增数据
	var res bool
	if permissionSaveVo.Type == model.PERMISSION_TYPE_GROUP {
		res, err = permissionService.SavePermissionGroup(permissionSaveVo)
	} else {
		res, err = permissionService.SavePermissionItem(permissionSaveVo)
	}
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	// 重新加载安全配置
	middleware.ClearPolicyData()

	middleware.SetData(c, res)
}

func AdminPermissionRemove(c *gin.Context) {
	// 获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	res, err = permissionService.RemovePermissionById(id)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	// 重新加载安全配置
	middleware.ClearPolicyData()
	
	middleware.SetData(c, res)
}

func AdminPermissionUpdate(c *gin.Context) {
	// 绑定并检查参数
	var permissionUpdateVo request.PermissionUpdateVo
	var err error
	err = c.ShouldBind(&permissionUpdateVo)
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	permissionUpdateVo.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.ReportError(c, response.PARAMETER_ILLEGAL, err)
	}
	// 更新记录
	var res bool
	if permissionUpdateVo.Type == model.PERMISSION_TYPE_GROUP {
		res, err = permissionService.UpdatePermissionGroup(permissionUpdateVo)
	} else {
		res, err = permissionService.UpdatePermissionItem(permissionUpdateVo)
	}
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}
	
	// 重新加载安全配置
	middleware.ClearPolicyData()

	middleware.SetData(c, res)
}

func AdminPermissionUpdateDeleted(c *gin.Context) {
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
	res, err = permissionService.UpdateDeleted(deletedVo)
	if err != nil {
		middleware.ReportError(c, response.SQL_FAILED, err)
	}

	// 重新加载安全配置
	middleware.ClearPolicyData()
	
	middleware.SetData(c, res)
}