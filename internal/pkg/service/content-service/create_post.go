package content_service

import (
	"context"
	"fmt"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// CreatePost создает новый пост
func (s *Service) CreatePost(ctx context.Context, post content_model.Post) (*content_model.Post, error) {
	createdPost, err := s.storage.CreatePost(ctx, post.ToInsertMap())
	if err != nil {
		return nil, fmt.Errorf("s.storage.CreatePost: %w", err)
	}

	return createdPost, nil
}
