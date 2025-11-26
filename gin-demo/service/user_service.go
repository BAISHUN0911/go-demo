package service

import (
	"errors"

	"gin-demo/model"
	"gin-demo/repository"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound  = errors.New("用户不存在")
	ErrUserNotActive = errors.New("用户未激活")
	ErrUnknownRole   = errors.New("未知角色")
)

func Login(username string) (*model.User, string, error) {
	user, err := repository.GetUserByUsername(username)
	if err == gorm.ErrRecordNotFound {
		return nil, "", ErrUserNotFound
	} else if err != nil {
		return nil, "", err
	}

	if !user.Active {
		return nil, "", ErrUserNotActive
	}

	switch user.Role {
	case "admin":
		return user, "欢迎管理员", nil
	case "user":
		return user, "欢迎普通用户", nil
	default:
		return nil, "", ErrUnknownRole
	}
}
