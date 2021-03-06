package Routes

import (
	"github.com/gin-gonic/gin"
	"go-study/App/Api/User"
	"go-study/App/Index"
	"go-study/MiddleWare"
	"net/http"
)

//
// Web
// @Description: 默认路由
// @param r
// @return *gin.Engine
//
func Web(r *gin.Engine) {
	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	//关联控制器
	r.GET("/index/home", Index.Home)

	//路由组
	group := r.Group("/group")
	{
		group.GET("/home1", Index.Home)
	}

	//中间件
	r.GET("/middleware", MiddleWare.Middle(), Index.Middle)

	//用户信息
	user := r.Group("/user")
	{
		user.GET("/userInfo", User.GetUserInfo)
		user.GET("/add", User.Add)
	}

	//性能测试 加入用户数据
	r.GET("/test/addUser999", User.AddUser999)

}
