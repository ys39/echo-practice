package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"echo-practice/models"

	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

// PostController は投稿に関連するコントローラ
type PostController struct {
	Repo models.PostRepository
}

// 投稿の詳細取得のコントローラ
// ※他Packageで使用するため、関数名の先頭を大文字にしている
func (p *PostController) GetPostDetail(c echo.Context) error {
	// mutexを使用して排他制御を行う
	lock.Lock()
	defer lock.Unlock() // 関数を抜ける際にmutexを解放する

	// パスパラメータの取得
	idStr := c.Param("id")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID: The 'id' parameter must be a positive integer.")
	}

	// データベースまたはリポジトリから投稿を取得
	// PostRepositoryインタフェースを実装した構造体よりメソッドを呼び出す
	post, err := p.Repo.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "The post with ID "+idStr+" was not found.")
	}

	return c.JSON(http.StatusOK, post)
}
