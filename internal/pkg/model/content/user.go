package content_model

import (
	desc "gqlgen-playground/internal/pb/content/v1"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/guregu/null.v4/zero"
)

// User структура для работы с пользователем
type User struct {
	Id        string      `db:"id" json:"id"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt zero.Time   `db:"deleted_at" json:"deleted_at"`
	Username  string      `db:"username" json:"username"`
	Name      string      `db:"name" json:"name"`
	Email     zero.String `db:"email" json:"email"`
	Phone     zero.String `db:"phone" json:"phone"`
}

// ToInsertMap возвращает мапу для вставки записи в БД
func (u *User) ToInsertMap() map[string]interface{} {
	if u == nil {
		return nil
	}

	return map[string]interface{}{
		"id":         uuid.NewString(),
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"username":   u.Username,
		"name":       u.Name,
		"email":      u.Email,
		"phone":      u.Phone,
	}
}

// ToProto конвертирует модель к прото модели
func (u *User) ToProto() *desc.User {
	if u == nil {
		return nil
	}

	result := &desc.User{
		Id:        u.Id,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email.Ptr(),
		Phone:     u.Phone.Ptr(),
	}

	if u.DeletedAt.Valid {
		result.DeletedAt = timestamppb.New(u.DeletedAt.Time)
	}

	return result
}
