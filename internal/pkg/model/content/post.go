package content_model

import (
	desc "gqlgen-playground/internal/pb/content/v1"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/guregu/null.v4/zero"
)

// Post структура для работы с постом
type Post struct {
	Id        string    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt zero.Time `db:"deleted_at" json:"deleted_at"`
	AuthorID  string    `db:"author_id" json:"author_id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
}

// ToInsertMap возвращает мапу для вставки записи в БД
func (u *Post) ToInsertMap() map[string]interface{} {
	if u == nil {
		return nil
	}

	return map[string]interface{}{
		"id":         uuid.NewString(),
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"author_id":  u.AuthorID,
		"title":      u.Title,
		"body":       u.Body,
	}
}

// ToProto конвертирует модель к прото модели
func (u *Post) ToProto() *desc.Post {
	if u == nil {
		return nil
	}

	result := &desc.Post{
		Id:        u.Id,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		AuthorId:  u.AuthorID,
		Title:     u.Title,
		Body:      u.Body,
	}

	if u.DeletedAt.Valid {
		result.DeletedAt = timestamppb.New(u.DeletedAt.Time)
	}

	return result
}
