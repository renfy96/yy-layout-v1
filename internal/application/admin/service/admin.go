package service

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/application/admin/params"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/application"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg"
	"github.com/renfy96/yy-layout-v1/pkg/helper/resp"
	"net/http"
)

type Service struct {
	query   application.QueryService
	command application.CommandService
}

func NewService(query application.QueryService, svc application.CommandService) Service {
	return Service{query: query, command: svc}
}

func (svc Service) Login(ctx *gin.Context) {
	var (
		err      error
		req      params.LoginReq
		adminAgg agg.Admin
	)
	if err = ctx.ShouldBind(&req); err != nil {
		resp.HandleParamsErr(ctx)
		return
	}
	if adminAgg, err = svc.query.GetAdminByAccAndPwd(ctx, req.Account, req.Password); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, adminAgg.Value().ID)
}
