package middle_ware

import "github.com/gin-gonic/gin"

type Au struct {
	Token string `json:"token"`
}
type Login struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var login Login
		err := ctx.ShouldBindJSON(&login)
		if err != nil {
			ctx.JSON(500, "参数错误")
			return
		}
		if login.Username != "admin" {
			ctx.JSON(401, "用户名错误")
		}
		if login.Password != "admin" {
			ctx.JSON(401, "密码错误")
		}
		login.UserID = 1
		token := SetJWT(ctx, login)
		ctx.JSON(200, map[string]string{"token": token, "msg": "登录成功"})
		return
	}
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		CheckJWT(ctx, token)
		return
	}
}
