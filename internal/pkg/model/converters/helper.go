package converters

import (
	"fmt"
	contentv1 "gqlgen-playground/internal/pb/content/v1"
	"gqlgen-playground/internal/pkg/model"
	"gqlgen-playground/internal/pkg/model/content"
	"time"

	"github.com/google/uuid"
)

// ==== Маппинг pb → GraphQL-модели ====

func PbUserToModel(u *contentv1.User) (*content.User, error) {
	if u == nil {
		return nil, nil
	}
	id, err := uuid.Parse(u.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid user id %q: %w", u.GetId(), err)
	}

	return &content.User{
		ID:   id,
		Name: u.GetName(),
	}, nil
}

func PbPostToModel(p *contentv1.Post) (*model.Post, error) {
	if p == nil {
		return nil, nil
	}
	id, err := uuid.Parse(p.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid post id %q: %w", p.GetId(), err)
	}

	return &model.Post{
		ID:      id,
		Title:   p.GetTitle(),
		Content: p.GetBody(), // body из gRPC → content в GraphQL
		// Author и Comments пока не заполняем — для N+1 будем делать отдельные резолверы.
	}, nil
}

func PbCommentToModel(c *contentv1.Comment) (*model.Comment, error) {
	if c == nil {
		return nil, nil
	}
	id, err := uuid.Parse(c.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid comment id %q: %w", c.GetId(), err)
	}

	var createdAt time.Time
	if ts := c.GetCreatedAt(); ts != nil {
		createdAt = ts.AsTime()
	}

	return &model.Comment{
		ID:        id,
		Content:   c.GetBody(),
		CreatedAt: createdAt,
		// Author, Replies пока не трогаем.
	}, nil
}
