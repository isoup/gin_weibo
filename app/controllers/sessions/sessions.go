package sessions

import (
	"gin_weibo/app/controllers"
	"gin_weibo/pkg/flash"

	userRequest "gin_weibo/app/requests/user"

	"github.com/gin-gonic/gin"
)

// Create 登录界面
func Create(c *gin.Context) {
	controllers.Render(c, "sessions/create.html", gin.H{})
}

// Store 登录 (创建新会话)
func Store(c *gin.Context) {
	userLoginForm := &userRequest.UserLoginForm{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	user, errors := userLoginForm.ValidateAndGetUser(c)

	if len(errors) != 0 || user == nil {
		flash.SaveValidateMessage(c, errors)
		controllers.Redirect(c, "/login")
		return
	}

	flash.NewSuccessFlash(c, "欢迎回来！")
	controllers.RedirectToUserShowPage(c, user)
}

// Destroy 登出 (销毁会话)
func Destroy(c *gin.Context) {

}
