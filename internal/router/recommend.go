package router

import (
	"github.com/gin-gonic/gin"
	"github.com/whj1990/bff-grpc/internal/handler"
	"github.com/whj1990/go-common/constant"
)

type RouteHandler struct {
	clientHandler handler.ClientHandler
}

func (r *RouteHandler) SetRouter(app *gin.Engine) {
	group := app.Group(constant.MineRouterGroup)
	group.GET("/review/project/list", r.clientHandler.ReviewProjectList())
}
