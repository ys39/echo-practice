/*
* テストの実行
 */
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"echo-practice/repositories/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	post1JSON      = `{"id":1,"title":"投稿1","content":"サンプル投稿1"}`
	createPostJSON = `{"id":11,"title":"create title","content":"create content"}`
	updatePostJSON = `{"id":1,"title":"update title","content":"update content"}`
)

func TestSuccessGetPostDetail(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/v1/api/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/detail/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.NoError(t, postController.GetPostDetail(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, post1JSON, rec.Body.String())
	}
}

/*
func TestFailGetPostDetail(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/v1/api/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/detail/:id")
	c.SetParamNames("id")
	c.SetParamValues("100")

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.Error(t, postController.GetPostDetail(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
*/

func TestSuccessGetPosts(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/v1/api/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list")

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.NoError(t, postController.GetPosts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSuccessCreatePost(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストボディを作成
	postData := map[string]string{
		"title":   "create title",
		"content": "create content",
	}
	jsonData, _ := json.Marshal(postData)

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodPost, "/v1/api/", bytes.NewBuffer(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/create")

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.NoError(t, postController.CreatePost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, createPostJSON, rec.Body.String())
	}
}

func TestSuccessUpdatePost(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストボディを作成
	postData := map[string]string{
		"title":   "update title",
		"content": "update content",
	}
	jsonData, _ := json.Marshal(postData)

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodPut, "/v1/api/", bytes.NewBuffer(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/update/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.NoError(t, postController.UpdatePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, updatePostJSON, rec.Body.String())
	}
}

func TestSuccessDeletePost(t *testing.T) {
	e := echo.New()

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodDelete, "/v1/api/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/delete/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))

	// モックのリポジトリを作成
	repo := &mock.MockdbPostRepository{
		Posts: mock.Posts,
	}

	// コントローラを初期化
	// Repositoryをコントローラに注入
	postController := &PostController{Repo: repo}

	// テスト用のリクエストを実行
	if assert.NoError(t, postController.DeletePost(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
