/*
* PostRepository インターフェース
 */

package models

// PostRepository インターフェース
type PostRepository interface {
	FindByID(id int) (*Post, error)           // IDに基づいて投稿を取得
	FindAll() ([]Post, error)                 // 全ての投稿を取得
	Create(post *Post) (*Post, error)         // 投稿を作成
	Update(id int, post *Post) (*Post, error) // 投稿を更新
	Delete(id int) error                      // 投稿を削除
}
