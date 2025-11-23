package content_service

import (
	"context"
	"fmt"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// CreateUser создает нового пользователя
func (s *Service) CreateUser(ctx context.Context, user content_model.User) (*content_model.User, error) {
	createdUser, err := s.storage.CreateUser(ctx, user.ToInsertMap())
	if err != nil {
		return nil, fmt.Errorf("s.storage.CreateUser: %w", err)
	}

	return createdUser, nil
}
