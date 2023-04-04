package service

import (
	"blog-server-go/model"
	"blog-server-go/model/request"
	"blog-server-go/model/response"
	"blog-server-go/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {}

//
// GetUserByUsername
//  @Description: 通过用户名获取用户
//  @receiver a
//  @param username
//  @return *model.User
//  @return error
//
func (a *UserService) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	// 没有查询到记录则返回空值
	var count int64
	err := db.Table("user").Where("username = ?", username).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	// 查询存在的记录
	err = db.Preload("RoleList").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//
// GetUserInfoDtoById
//  @Description: 通过id获取用户信息
//  @receiver a
//  @param id
//  @return *response.UserInfoDto
//  @return error
//
func (a *UserService) GetUserInfoDtoById(id int) (*response.UserInfoDto, error) {
	var user model.User
	err := db.Preload("RoleList").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user.CopyToUserInfoDto()
}

func (a *UserService) CountUser() (int, error) {
	var count int
	err := db.Model(&model.User{}).Count(&count).Error
	return count, err
}

//
// UserRegister
//  @Description: 用户注册
//  @receiver a
//  @param registerVo
//  @return bool
//  @return error
//
func (a *UserService) UserRegister(registerVo request.UserRegisterVo) (bool, error) {
	// 构造用户
	var user model.User
	err := utils.CopyFields(&user, registerVo)
	if err != nil {
		return false, err
	}
	user.Username = registerVo.Phone
	// 密码加密
	var passwordCrypt []byte
	passwordCrypt, err = bcrypt.GenerateFromPassword([]byte(registerVo.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user.Password = string(passwordCrypt)
	// 新增用户，成功后自动返回id
	tx := db.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 绑定角色
	err = tx.Model(&user).Association("RoleList").Append(&model.Role{
		Model: model.Model{ ID: int(model.ROLE_USER_ID) },
	}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

//
// UpdateUser
//  @Description: 修改用户信息，一般用于用户自身操作
//  @receiver a
//  @param userUpdateVo
//  @return bool
//  @return error
//
func (a *UserService) UpdateUser(userUpdateVo request.UserUpdateVo) (bool, error) {
	// 检查手机号是否重合
	if userUpdateVo.Phone != "" {
		var count int
		db.Where("phone = ? and id <> ?", userUpdateVo.Phone, userUpdateVo.ID).Count(&count)
		if count > 0 {
			return false, errors.New("该手机号已被他人注册")
		}
	}
	var user model.User
	err := utils.CopyFields(&user, userUpdateVo)
	if err != nil {
		return false, err
	}

	err = db.Model(&user).Updates(user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// UpdatePassword
//  @Description: 修改密码
//  @receiver a
//  @param userPasswordVo
//  @return bool
//  @return error
//
func (a *UserService) UpdatePassword(userPasswordVo request.UserPasswordVo) (bool, error) {
	// 检查原密码
	var user model.User
	err := db.Where("id = ?", userPasswordVo.ID).First(&user).Error
	if err != nil {
		return false, errors.New("用户不存在")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPasswordVo.Password)) != nil {
		return false, errors.New("原密码不正确")
	}
	// 修改密码
	var passwordCrypt []byte
	passwordCrypt, err = bcrypt.GenerateFromPassword([]byte(userPasswordVo.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	err = db.Model(&user).Update("password", string(passwordCrypt)).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

//
// UpdateDeleted
//  @Description: 更新用户禁用状态，操作完成后，建议将用户下线
//  @receiver a
//  @param modelDeletedVo
//  @return bool
//  @return error
//
func (a *UserService) UpdateDeleted(modelDeletedVo request.ModelDeletedVo) (bool, error) {
	var user model.User
	user.ID = modelDeletedVo.ID
	err := db.Model(&user).Update("deleted", modelDeletedVo.Deleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// RemoveUserById
//  @Description: 通过id删除用户
//  @receiver a
//  @param id
//  @return bool
//  @return error
//
func (a *UserService) RemoveUserById(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//
// UserAdminDtoPage
//  @Description: 分页获取后台用户
//  @receiver a
//  @param pageVo
//  @param userVo
//  @return *response.Page
//  @return error
//
func (a *UserService) UserAdminDtoPage(pageVo request.PageVo, userVo request.UserVo) (*response.Page, error) {
	// 构造查询条件
	limit := pageVo.Size
	offset := pageVo.Size * (pageVo.Current - 1)
	// 预加载角色列表
	tx := db.Preload("RoleList")
	// 动态拼接查询条件
	if userVo.Nickname != "" {
		tx = tx.Where("nickname LIKE ?", "%" + userVo.Nickname + "%")
	}
	if userVo.Phone != "" {
		tx = tx.Where("phone LIKE ?", "%" + userVo.Phone + "%")
	}
	// 查询分页信息
	var userList []*model.User
	var count int64
	err := tx.Limit(limit).Offset(offset).Find(&userList).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var userAdminList []*response.UserAdminDto
	for _, item := range userList {
		var userAdmin *response.UserAdminDto
		userAdmin, err = item.CopyToUserAdminDto()
		userAdminList = append(userAdminList, userAdmin)
	}
	// 构造分页
	page := response.Page{
		Current: pageVo.Current,
		Size:    int64(len(userAdminList)),
		Total:   count,
		Records: userAdminList,
	}
	return &page, nil
}

//
// UpdateUserAdminDto
//  @Description: 更新后台用户更新，仅支持更新昵称和角色列表，操作完成后，建议将用户下线
//  @receiver a
//  @return bool
//  @return error
//
func (a *UserService) UpdateUserAdminDto(userAdminUpdateVo request.UserAdminUpdateVo) (bool, error) {
	var user model.User
	err := utils.CopyFields(&user, userAdminUpdateVo)
	if err != nil {
		return false, err
	}
	tx := db.Begin()
	// 更新基本信息
	err = tx.Model(&user).Updates(user).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 更新角色关联信息
	if len(userAdminUpdateVo.RoleIdList) != 0 {
		// 删除原有角色关系
		err = tx.Model(&user).Association("RoleList").Clear().Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
		// 插入新的角色关系
		var roleList []model.Role
		for _, roleId := range userAdminUpdateVo.RoleIdList {
			role := model.Role{Model: model.Model{ ID: roleId }}
			roleList = append(roleList, role)
		}
		err = tx.Model(&user).Association("RoleList").Append(roleList).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

