package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg"
)

type AdminRepository interface {
	GetAdminByAccount(ctx *gin.Context, name string) agg.Admin
	GetAdminById(ctx *gin.Context, id uint) agg.Admin
	Save(ctx *gin.Context, admin agg.Admin) (err error)
}
