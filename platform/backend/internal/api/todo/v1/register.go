package api_todo_v1

import (
	"api/go-sdk/todo/v1/todoV1connect"
	"net/http"
	gendbtodo "platform/backend/internal/gendb/todo"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Store gendbtodo.Querier
}

func Register(apis *gin.RouterGroup, deps Dependencies) {
	service := newTodoService(deps.Store)
	servicePath, serviceHandler := todoV1connect.NewTodoServiceHandler(service)

	apis.POST(servicePath+"*rpc", gin.WrapH(http.StripPrefix("/api", serviceHandler)))
}
