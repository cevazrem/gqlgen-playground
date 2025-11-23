package content_storage

import (
	"context"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"

	sq "github.com/Masterminds/squirrel"
)

// CreateComment создать комментарий
func (s *Storage) CreateComment(ctx context.Context, data map[string]interface{}) (*content_model.Comment, error) {
	query := s.builder().
		Insert(CommentTable).
		SetMap(data).
		Suffix(returnAllCommentsTabledFields)

	comment := new(content_model.Comment)
	if pgErr := s.getx(ctx, comment, query); pgErr != nil {
		return nil, handlePgError(pgErr)
	}

	return comment, nil
}

// ListComments получить список комментариев из БД
func (s *Storage) ListComments(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.Comment, int64, error) {
	qb := s.builder().
		Select(commentsTabledFields).
		From(CommentTable).
		Where(whereCond).
		OrderBy("created_at DESC").
		Limit(pagination.Limit()).
		Offset(pagination.Offset())

	var comments []*content_model.Comment
	if err := s.selectx(ctx, &comments, qb); err != nil {
		return nil, 0, handlePgError(err)
	}

	cntb := s.builder().
		Select("count (*) as total").
		From(CommentTable).
		Where(whereCond)

	var total int64
	if err := s.getx(ctx, &total, cntb); err != nil {
		return nil, 0, handlePgError(err)
	}

	return comments, total, nil
}
