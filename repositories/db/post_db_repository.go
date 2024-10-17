/*
* DBのリポジトリ実装
 */

package db

import (
	"echo-practice/models"
	"errors"
)

// DbPostRepository はPostRepositoryのDB実装
type DbPostRepository struct {
	Posts []models.Post
}

// FindByID はIDに基づいて投稿を取得
func (r *DbPostRepository) FindByID(id int) (*models.Post, error) {
	// DBから取得したデータを返す
	return nil, errors.New("post not found")
}
