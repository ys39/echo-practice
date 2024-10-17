/*
* コントローラ層
 */

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

// 投稿一覧取得のコントローラ
func (p *PostController) GetPosts(c echo.Context) error {
	// mutexを使用して排他制御を行う
	lock.Lock()
	defer lock.Unlock() // 関数を抜ける際にmutexを解放する

	// データベースまたはリポジトリから投稿を取得
	// PostRepositoryインタフェースを実装した構造体よりメソッドを呼び出す
	posts, err := p.Repo.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No posts were found.")
	}

	return c.JSON(http.StatusOK, posts)
}

// 投稿の作成のコントローラ
func (p *PostController) CreatePost(c echo.Context) error {
	// mutexを使用して排他制御を行う
	lock.Lock()
	defer lock.Unlock() // 関数を抜ける際にmutexを解放する

	// リクエストボディの取得
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body.")
	}

	// データベースまたはリポジトリに投稿を作成
	// PostRepositoryインタフェースを実装した構造体よりメソッドを呼び出す
	createdPost, err := p.Repo.Create(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create a post.")
	}

	return c.JSON(http.StatusCreated, createdPost)
}

// 投稿の更新のコントローラ
func (p *PostController) UpdatePost(c echo.Context) error {
	// mutexを使用して排他制御を行う
	lock.Lock()
	defer lock.Unlock() // 関数を抜ける際にmutexを解放する

	// パスパラメータの取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID: The 'id' parameter must be a positive integer.")
	}

	// リクエストボディの取得
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body.")
	}

	// データベースまたはリポジトリに投稿を更新
	// PostRepositoryインタフェースを実装した構造体よりメソッドを呼び出す
	updatedPost, err := p.Repo.Update(id, post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update the post.")
	}

	return c.JSON(http.StatusOK, updatedPost)
}

// 投稿の削除のコントローラ
func (p *PostController) DeletePost(c echo.Context) error {
	// mutexを使用して排他制御を行う
	lock.Lock()
	defer lock.Unlock() // 関数を抜ける際にmutexを解放する

	// パスパラメータの取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID: The 'id' parameter must be a positive integer.")
	}

	// データベースまたはリポジトリから投稿を削除
	// PostRepositoryインタフェースを実装した構造体よりメソッドを呼び出す
	err = p.Repo.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete the post.")
	}

	return c.NoContent(http.StatusNoContent)
}
