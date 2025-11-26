package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-demo/src/common/model"
)

func ConditionDemo() {
	age := 20

	// Go 条件不需要括号 ()
	if age < 18 {
		fmt.Println("未成年")
	} else if age >= 18 && age < 60 {
		fmt.Println("成年人")
	} else {
		fmt.Println("老年人")
	}

	// Go 可以在 if 中先声明变量
	if isAdult := age >= 18; isAdult {
		fmt.Println("可以进入成人区")
	}

	user := getUser("bob")
	if user != nil { // user != nil 表示 user 不为 nil
		// 将 Go 对象转换为 JSON 字节数组
		userJSON, err := json.Marshal(user)
		if err != nil {
			fmt.Println("转换JSON出错", err)
		} else {
			fmt.Println("用户JSON:", string(userJSON))
		}
	}
	err := loginHandler(user)
	if err != nil {
		fmt.Println("登录失败:", err)
	}

	// switch
	switch "1" {
	case "1":
		println("1")
	case "2":
		println("2")
	default:
		println("default")
	}
}

// 模拟获取用户（真实场景可能从数据库查）
func getUser(username string) *model.User {
	switch username {
	case "alice":
		return &model.User{Username: "alice", Role: "admin", Active: true}
	case "bob":
		return &model.User{Username: "bob", Role: "user1", Active: true}
	}
	return nil
}

// 用户登录处理
func loginHandler(user *model.User) error {
	// 在 if 条件中声明变量
	if user == nil {
		fmt.Println("用户不存在")
		return nil
	} else if !user.Active {
		fmt.Println("用户未激活，请联系管理员")
		return nil
	} else if user.Role == "admin" {
		fmt.Println("欢迎管理员", user.Username)
		return nil
	} else if user.Role == "user" {
		fmt.Println("欢迎普通用户", user.Username)
		return nil
	} else {
		fmt.Println("未知角色")
		return errors.New("未知角色")
	}
}
