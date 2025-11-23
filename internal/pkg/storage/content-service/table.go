package content_storage

import (
	"fmt"
	content_model "gqlgen-playground/internal/pkg/model/content"
	"reflect"
)

const returningSuffix = "RETURNING"

const (
	// UsersTable наименование таблицы пользователей в БД
	UsersTable = "users"
	// PostTable наименование таблицы постов в БД
	PostTable = "posts"
	// CommentTable наименование таблицы комментариев в БД
	CommentTable = "comments"
)

var (
	usersTabledFields    = allFields(content_model.User{})
	postsTabledFields    = allFields(content_model.Post{})
	commentsTabledFields = allFields(content_model.Comment{})
)

var (
	returnAllUsersTabledFields    = fmt.Sprintf("%s %s", returningSuffix, usersTabledFields)
	returnAllPostsTabledFields    = fmt.Sprintf("%s %s", returningSuffix, postsTabledFields)
	returnAllCommentsTabledFields = fmt.Sprintf("%s %s", returningSuffix, commentsTabledFields)
)

// allFields возвращает список всех полей структуры, которые есть в БД
func allFields(data interface{}) string {
	var s string
	r := reflect.TypeOf(data)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i).Tag.Get("db")
		if field != "" {
			s += field + ","
		}
	}
	return s[:len(s)-1]
}
