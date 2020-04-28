package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	myjwt "singo/middleware/jwt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return generateToken(c, user)

}

type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) serializer.Response {
	j := &myjwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		user.Id,
		user.Nickname,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		return serializer.ParamErr("登录信息认证失败", nil)

	}

	return serializer.BuildUserResponse(user, token)
}
