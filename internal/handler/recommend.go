package handler

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/whj1990/bff-grpc/internal/handler/stru"
	"github.com/whj1990/bff-grpc/rpc"
	"github.com/whj1990/go-core/util"
	"github.com/whj1990/mine-grrpc/pbs"
)

type ClientHandler interface {
	ReviewProjectList() gin.HandlerFunc
	StreamClientServer() gin.HandlerFunc
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
		r.client = rpc.NewGrpcClient()
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
func (r *clientHandler) StreamClientServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r.client = rpc.NewGrpcClient()
		res, err := r.client.StreamClientServer(util.GetRPCContext(ctx))
		for i := 0; i < 20; i++ {
			err := res.Send(&pbs.ParamId{
				Num: int32(i),
			})
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("stream request err: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		r, err := res.CloseAndRecv()
		fmt.Println(r)
		if err != nil {
			stru.HandleErrorResponse(ctx, err)
			return
		}
		stru.HandleSuccessDataResponse(ctx, res)
	}
}
