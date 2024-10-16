package main

import (
	"echo-practice/errors"
	"echo-practice/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// カスタムエラーハンドラを登録
	e.HTTPErrorHandler = errors.CustomHTTPErrorHandler

	// ルーティング
	routers.SetupRouter(e)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
