package content_storage

import (
	"context"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"

	sq "github.com/Masterminds/squirrel"
)

// CreatePost создать пост
func (s *Storage) CreatePost(ctx context.Context, data map[string]interface{}) (*content_model.Post, error) {
	query := s.builder().
		Insert(PostTable).
		SetMap(data).
		Suffix(returnAllPostsTabledFields)

	post := new(content_model.Post)
	if pgErr := s.getx(ctx, post, query); pgErr != nil {
		return nil, handlePgError(pgErr)
	}

	return post, nil
}

// GetPost получить пост из БД
func (s *Storage) GetPost(ctx context.Context, whereCond sq.Sqlizer) (*content_model.Post, error) {
	qb := s.builder().
		Select(postsTabledFields).
		From(PostTable).
		Where(whereCond)

	post := new(content_model.Post)
	if err := s.getx(ctx, post, qb); err != nil {
		return nil, handlePgError(err)
	}

	return post, nil
}

// ListPosts получить список постов из БД
func (s *Storage) ListPosts(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.Post, int64, error) {
	qb := s.builder().
		Select(postsTabledFields).
		From(PostTable).
		Where(whereCond).
		OrderBy("created_at DESC").
		Limit(pagination.Limit()).
		Offset(pagination.Offset())

	var posts []*content_model.Post
	if err := s.selectx(ctx, &posts, qb); err != nil {
		return nil, 0, handlePgError(err)
	}

	cntb := s.builder().
		Select("count (*) as total").
		From(PostTable).
		Where(whereCond)

	var total int64
	if err := s.getx(ctx, &total, cntb); err != nil {
		return nil, 0, handlePgError(err)
	}

	return posts, total, nil
}
