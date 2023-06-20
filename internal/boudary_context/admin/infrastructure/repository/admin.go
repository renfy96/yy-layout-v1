package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/boudary_context/admin/domain/agg"
	"github.com/renfy96/yy-layout-v1/internal/boudary_context/admin/domain/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repository.AdminRepository {
	return repositoryImpl{
		db: db,
	}
}

func (r repositoryImpl) GetAdminByAccount(ctx *gin.Context, name string) agg.Admin {
	//TODO implement me
	panic("implement me")
}

func (r repositoryImpl) GetAdminById(ctx *gin.Context, id uint) agg.Admin {
	//TODO implement me
	panic("implement me")
}

func (r repositoryImpl) Save(ctx *gin.Context, admin agg.Admin) {
	//TODO implement me
	panic("implement me")
}
