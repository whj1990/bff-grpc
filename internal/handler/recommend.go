package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/whj1990/bff-grpc/internal/handler/stru"
	"github.com/whj1990/go-core/util"
	"github.com/whj1990/mine-grrpc/pbs"
)

type ClientHandler interface {
	ReviewProjectList() gin.HandlerFunc
}
type clientHandler struct {
	client pbs.HandleServerClient
}

// @Summary	获取项目列表
// @Tags		review
// @Produce	json
// @Param		query			query		stru.ReviewProjectListParams						false	"query"
// @Success	200				{object}	stru.Response{data=pbs.ReviewProjectListResponse}	"ok"
// @Router		/mine/review/project/list [GET]
func (r *clientHandler) ReviewProjectList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.ReviewProjectListParams
		if err := ctx.ShouldBindQuery(&req); err != nil {
			stru.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.client.ReviewProjectList(util.GetRPCContext(ctx), &pbs.ReviewProjectListParams{
			PageNum:    req.PageNum,
			PageSize:   req.PageSize,
			Id:         req.Id,
			Name:       req.Name,
			ModeCode:   req.ModeCode,
			Status:     req.Status,
			ShowStatus: req.ShowStatus,
		})
		if err != nil {
			stru.HandleErrorResponse(ctx, err)
			return
		}
		stru.HandleSuccessDataResponse(ctx, res)
	}
}
