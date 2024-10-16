/*
* モデル層
* データベースとのやり取りを行う(簡易的にサンプルデータを使用)
* データ構造体を定義する
 */

package models

// 投稿データ構造体を定義する
// 「タグ」機能を用いることで、構造体のフィールドとJSONデータの間で変換を行う
type (
	Post struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
)
