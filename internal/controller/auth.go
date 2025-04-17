package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"sphinx/internal/models/qo"
	"sphinx/internal/models/vo"
	"sphinx/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// Login 检查手机号是否存在
func Login(c *gin.Context) {
	// 解析JSON请求体
	var account qo.AccountRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		zap.L().Error("解析请求失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("解析请求失败").Finished())
		return
	}

	// 正则检查手机号格式
	matched, err := regexp.MatchString("^1[3456789]\\d{9}$", account.Phone)
	if err != nil {
		zap.L().Error("校验手机号失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("校验手机号失败").Finished())
		return
	}
	if !matched {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("手机号格式错误").Finished())
		return
	}

	// 如果有密码字段，则进行登录校验
	if account.Password != "" {
		// TODO: 实现密码验证逻辑
		// 模拟登录成功，返回重定向URL
		c.JSON(http.StatusOK, vo.Builder().Ok().Data(map[string]string{
			"redirectUrl": "http://example.com/dashboard",
		}).Msg("登录成功").Finished())
		return
	}

	// 检查手机号是否存在
	exists, err := service.CheckPhoneExists(account.Phone)
	fmt.Printf("exists: %v, err: %v", exists, err)
	if err != nil {
		zap.L().Error("检查手机号失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("检查手机号失败").Finished())
		return
	}
	if !exists {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.AccountNotExist).Finished())
		return
	}

	c.JSON(http.StatusOK, vo.Builder().Ok().Msg("校验通过").Finished())
}

// Register 检查密码是否存在
func Register(c *gin.Context) {
	// 解析JSON请求体
	var account qo.AccountRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		zap.L().Error("解析请求失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("解析请求失败").Finished())
		return
	}

	// 正则检查手机号格式
	matched, err := regexp.MatchString("^1[3456789]\\d{9}$", account.Phone)
	if err != nil {
		zap.L().Error("校验手机号失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("校验手机号失败").Finished())
		return
	}
	if !matched {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("手机号格式错误").Finished())
		return
	}

	// 检查手机号是否存在
	exists, err := service.CheckPhoneExists(account.Phone)
	if err != nil {
		zap.L().Error("检查手机号失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("检查手机号失败").Finished())
		return
	}
	if exists {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.AccountAlreadyExist).Finished())
		return
	}
	// 校验密码
	if account.Password == "" {
		c.JSON(http.StatusOK, vo.Builder().Ok().Msg("输入密码后点击注册按钮").Finished())
		return
	}
	if len(account.Password) < 6 || len(account.Password) > 16 {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("密码长度不能小于6位或大于16位").Finished())
		return
	}

	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("密码哈希失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("密码哈希失败").Finished())
		return
	}

	// 保存密码哈希
	err = service.SaveAccount(account.Phone, string(hashedPassword))
	if err != nil {
		zap.L().Error("保存密码失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("保存密码失败").Finished())
		return
	}

	c.JSON(200, vo.Builder().Ok().Msg("success").Finished())
}
