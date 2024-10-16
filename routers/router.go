/*
* ルーティングの設定を行う
 */

package routers

import (
	"echo-practice/controllers"
	"echo-practice/repositories/mock"

	"github.com/labstack/echo/v4"
)

// ルーティングの設定
func SetupRouter(e *echo.Echo) {

	// インメモリのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	postController := &controllers.PostController{Repo: repo}

	// /v1/api に関連するエンドポイントをグループ化
	api := e.Group("/v1/api")
	api.GET("/detail/:id", postController.GetPostDetail)
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
