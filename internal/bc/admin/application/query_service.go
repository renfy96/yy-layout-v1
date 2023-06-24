package application

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/repository"
)

type QueryService struct {
	adminRepo repository.AdminRepository
}

func NewQueryService(repo repository.AdminRepository) QueryService {
	return QueryService{adminRepo: repo}
}

func (svc QueryService) GetAdminByAccAndPwd(ctx *gin.Context, account, pwd string) (adminAgg agg.Admin, err error) {
	adminAgg = svc.adminRepo.GetAdminByAccount(ctx, account)
	if !adminAgg.IsExist() {
		err = errors.New("账号或密码错误")
		return
	}
	if !adminAgg.CheckPassword(pwd) {
		err = errors.New("账号或密码错误")
		return
	}
	return
}
