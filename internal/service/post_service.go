package service

import (
	"context"
)

// PostService define as operações de negócio para Post
type PostService interface {
	CreatePost(ctx context.Context, title, content, blogID, authorID string) error
	GetPost(ctx context.Context, id string) error
	ListPosts(ctx context.Context, page, limit int, blogID string) error
	UpdatePost(ctx context.Context, id, title, content string) error
	DeletePost(ctx context.Context, id string) error
}
