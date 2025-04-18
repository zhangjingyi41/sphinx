package controller

import (
	"crypto/sha256" // 新增
	"encoding/hex"  // 新增
	"net/http"
	"regexp"
	"sphinx/internal/models/qo"
	"sphinx/internal/models/vo"
	"sphinx/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		// 校验手机账号和密码
		result, err := service.CheckAccountPassword(account.Phone, account.Password)
		if err != nil {
			zap.L().Error("校验账号密码失败", zap.Error(err))
			c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("校验账号密码失败").Finished())
			return
		}
		if !result {
			c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.PasswordIncorrect).Finished())
			return
		}
		// TODO: 实现密码验证逻辑
		// （废弃）模拟登录成功，返回重定向URL
		// 账号密码登录不需要重定向，直接返回登录成功，因为能用账号密码登录的都是要登录sphinx后台
		// c.JSON(http.StatusOK, vo.Builder().Ok().Data(map[string]string{
		// 	"redirectUrl": "http://121.199.164.40:17500",
		// }).Finished())
		// 生成token和refreshtoken
		token, refreshtoken, err := service.GenerateToken(account.Phone)
		if err != nil {
			zap.L().Error("生成token失败", zap.Error(err))
			c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("生成token失败").Finished())
			return
		}
		c.JSON(http.StatusOK, vo.Builder().Ok().Data(map[string]string{
			"token":        token,
			"refreshtoken": refreshtoken,
		}).Finished())

		// c.JSON(http.StatusOK, vo.Builder().Ok().Finished())
		return
	}

	// 检查手机号是否存在
	exists, err := service.CheckPhoneExists(account.Phone)
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
	hasher := sha256.New()
	hasher.Write([]byte(account.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	// 保存密码哈希
	err = service.SaveAccount(account.Phone, hashedPassword)
	if err != nil {
		zap.L().Error("保存密码失败", zap.Error(err))
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Reason("保存密码失败").Finished())
		return
	}

	c.JSON(200, vo.Builder().Ok().Msg("success").Finished())
}
