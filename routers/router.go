/*
* ルーティングの設定を行う
 */

package routers

import (
	"echo-practice/controllers"

	"github.com/labstack/echo/v4"
)

// ルーティングの設定
func SetupRouter(e *echo.Echo) {

	// /v1/api に関連するエンドポイントをグループ化
	api := e.Group("/v1/api")
	api.GET("/detail/:id", controllers.GetPostDetail)
	/*
		 {
			 //api.GET("/list", controllers.GetPosts)
			 api.GET("/detail/:id",)
			 //api.POST("/create", controllers.CreatePost)
			 //api.DELETE("/delete/:id", controllers.DeletePost)
			 //api.PUT("/update/:id", controllers.UpdatePost)
		 }
	*/
}
