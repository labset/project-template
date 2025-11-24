package api_todo_v1

import (
	"api/go-sdk/todo/v1/todoV1connect"
	gendbtodo "platform/backend/internal/gendb/todo"
)

type todoService struct {
	store gendbtodo.Querier
}

func newTodoService(store gendbtodo.Querier) todoV1connect.TodoServiceHandler {
	return &todoService{
		store: store,
	}
}
