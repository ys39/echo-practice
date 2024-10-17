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

	/*
		// モックのリポジトリを作成
		repo := &db.DbPostRepository{}
	*/

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &controllers.PostController{Repo: repo}

	// /v1/api に関連するエンドポイントをグループ化
	api := e.Group("/v1/api")
	api.GET("/detail/:id", postController.GetPostDetail)
	api.GET("/list", postController.GetPosts)
	api.POST("/create", postController.CreatePost)
	api.DELETE("/delete/:id", postController.DeletePost)
	api.PUT("/update/:id", postController.UpdatePost)
}
