/*
* モックDBのリポジトリ
 */

package mock

import (
	"echo-practice/models"
	"errors"
)

// MockdbPostRepository はPostRepositoryのMock実装
type MockdbPostRepository struct {
	Posts []models.Post
}

// FindByID はIDに基づいて投稿を取得
func (r *MockdbPostRepository) FindByID(id int) (*models.Post, error) {
	for _, post := range r.Posts {
		if post.ID == id {
			return &post, nil
		}
	}
	return nil, errors.New("post not found")
}

// FindAll は全ての投稿を取得
func (r *MockdbPostRepository) FindAll() ([]models.Post, error) {
	return r.Posts, nil
}

// Create は投稿を作成
func (r *MockdbPostRepository) Create(post *models.Post) (*models.Post, error) {
	var newPost models.Post
	newPost.ID = len(r.Posts) + 1
	r.Posts = append(r.Posts, newPost)
	return &newPost, nil
}

// Update は投稿を更新
func (r *MockdbPostRepository) Update(id int, post *models.Post) (*models.Post, error) {
	for i, p := range r.Posts {
		if p.ID == id {
			r.Posts[i].Title = post.Title
			r.Posts[i].Content = post.Content
			return post, nil
		}
	}
	return nil, errors.New("post not found")
}
