package api

import (
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
	myjwt "singo/middleware/jwt"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// // LoginResult 登录结果结构
// type LoginResult struct {
// 	Token string `json:"token"`
// 	model.User
// }

// Login 登录
func Login(c *gin.Context) {
	var loginReq model.LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := model.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

// // 生成令牌
// func generateToken(c *gin.Context, user model.User) {
// 	j := &myjwt.JWT{
// 		[]byte("newtrekWang"),
// 	}
// 	claims := myjwt.CustomClaims{
// 		user.Id,
// 		user.Name,
// 		user.Phone,
// 		jwtgo.StandardClaims{
// 			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
// 			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
// 			Issuer:    "newtrekWang",                   //签名的发行者
// 		},
// 	}

// 	token, err := j.CreateToken(claims)

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": -1,
// 			"msg":    err.Error(),
// 		})
// 		return
// 	}

// 	log.Println(token)

// 	data := LoginResult{
// 		User:  user,
// 		Token: token,
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 0,
// 		"msg":    "登录成功！",
// 		"data":   data,
// 	})
// 	return
// }

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
