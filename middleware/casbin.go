package middleware

import (
	"blog-server-go/model"
	"blog-server-go/model/response"
	"blog-server-go/service"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"strings"
)

var permissionService = service.GlobalService.PermissionService

// 全局执行器
var enforcer *casbin.Enforcer

func init() {
	en, err := casbin.NewEnforcer("conf/rbac_model.conf", "conf/policy.csv")
	if err != nil {
		panic(err)
	}
	enforcer = en
}

//
// Casbin
//  @Description: rbac权限控制
//  @return gin.HandlerFunc
//
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如权限信息被修改，则下次访问重新加载
		if len(enforcer.GetPolicy()) == 0 {
			err := LoadPolicyData()
			if err != nil {
				ReportError(c, response.ACCESS_FAILED, err)
			}
		}
		// 获取访问主体和访问资源
		obj := c.Request.URL.Path
		claims := GetClaims(c)
		sub := ""
		if claims != nil {
			sub = claims.Subject
		}
		// 匿名访问，配置的白名单可自由访问
		// 实名访问，将检查权限和可匿名访问状态
		res, err := enforcer.Enforce(sub, obj, "*")
		if err != nil {
			ReportError(c, response.ACCESS_FAILED, err)
		}
		if !res {
			ReportError(c, response.ACCESS_DENIED, nil)
		}
		// 授权通过
		c.Next()
	}
}
//
// LoadPolicyData
//  @Description: 加载权限信息
//  @return error
//
func LoadPolicyData() error {
	permissionRoleDtoList, err := permissionService.PermissionRoleDtoList()
	if err != nil {
		return err
	}
	// 整理权限策略，格式为 角色-资源-行为
	var policies [][]string
	for _, permissionRoleDto := range permissionRoleDtoList {
		if permissionRoleDto.Url == nil || *permissionRoleDto.Url == "" {
			return errors.New("不可加载空权限！")
		}
		obj := *permissionRoleDto.Url
		for _, roleOptionDto := range permissionRoleDto.RoleList {
			sub := roleOptionDto.Code
			// 将资源中存在的*替换为{id}，暂时仅处理单个*的情况
			// 使用*作为通配符是spring security的做法，在golang中更推荐{}风格
			if strings.Contains(obj, "*") {
				obj = strings.ReplaceAll(obj, "*", "{id}")
			}
			policies = append(policies, []string{sub, obj, "*"})
		}
		// 可匿名访问的资源，为其添加匿名权限
		if *permissionRoleDto.Anonymous == model.PERMISSION_ANONYMOUS_ENABLE {
			policies = append(policies, []string{"anonymous", obj, "*"})
		}
	}
	// 添加权限策略
	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		return err
	}
	return nil
}

//
// ClearPolicyData
//  @Description: 清除权限信息
//
func ClearPolicyData() {
	enforcer.ClearPolicy()
}

//
// AddUserGroup
//  @Description: 添加指定用户的角色信息
//  @param username
//  @param roles
//  @return error
//
func AddUserGroup(username string, roles []string) error {
	var groupPolicies [][]string
	for _, role := range roles {
		groupPolicies = append(groupPolicies, []string{username, role})
	}
	_, err := enforcer.AddGroupingPolicies(groupPolicies)
	if err != nil {
		return err
	}
	return nil
}

//
// ClearUserGroup
//  @Description: 清除指定用户的角色信息
//  @param username
//  @return error
//
func ClearUserGroup(username string) error {
	_, err := enforcer.RemoveFilteredGroupingPolicy(0, username)
	if err != nil {
		return err
	}
	return nil
}
