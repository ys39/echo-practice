/*
* モックDBのリポジトリ実装
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
	newPost.Title = post.Title
	newPost.Content = post.Content
	r.Posts = append(r.Posts, newPost)
	return &newPost, nil
}

// Update は投稿を更新
func (r *MockdbPostRepository) Update(id int, post *models.Post) (*models.Post, error) {
	for i, p := range r.Posts {
		if p.ID == id {
			// JSONの値を更新
			r.Posts[i].Title = post.Title
			r.Posts[i].Content = post.Content
			// 更新した投稿を返す
			var newPost models.Post
			newPost.ID = id
			newPost.Title = post.Title
			newPost.Content = post.Content
			return &newPost, nil
		}
	}
	return nil, errors.New("post not found")
}

// Delete は投稿を削除
func (r *MockdbPostRepository) Delete(id int) error {
	for i, post := range r.Posts {
		if post.ID == id {
			r.Posts = append(r.Posts[:i], r.Posts[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}
