package content_service

import (
	"context"
	"fmt"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// CreateComment создает новый комментарий
func (s *Service) CreateComment(ctx context.Context, post content_model.Comment) (*content_model.Comment, error) {
	createdComment, err := s.storage.CreateComment(ctx, post.ToInsertMap())
	if err != nil {
		return nil, fmt.Errorf("s.storage.CreateComment: %w", err)
	}

	return createdComment, nil
}
