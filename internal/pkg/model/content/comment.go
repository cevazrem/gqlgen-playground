package content_model

import (
	desc "gqlgen-playground/internal/pb/content/v1"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/guregu/null.v4/zero"
)

// Comment структура для работы с комментарием
type Comment struct {
	Id        string    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt zero.Time `db:"deleted_at" json:"deleted_at"`
	PostID    string    `db:"post_id" json:"post_id"`
	AuthorID  string    `db:"author_id" json:"author_id"`
	Body      string    `db:"body" json:"body"`
}

// ToInsertMap возвращает мапу для вставки записи в БД
func (u *Comment) ToInsertMap() map[string]interface{} {
	if u == nil {
		return nil
	}

	return map[string]interface{}{
		"id":         uuid.NewString(),
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"post_id":    u.PostID,
		"author_id":  u.AuthorID,
		"body":       u.Body,
	}
}

// ToProto конвертирует модель к прото модели
func (u *Comment) ToProto() *desc.Comment {
	if u == nil {
		return nil
	}

	result := &desc.Comment{
		Id:        u.Id,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		PostId:    u.PostID,
		AuthorId:  u.AuthorID,
		Body:      u.Body,
	}

	if u.DeletedAt.Valid {
		result.DeletedAt = timestamppb.New(u.DeletedAt.Time)
	}

	return result
}
