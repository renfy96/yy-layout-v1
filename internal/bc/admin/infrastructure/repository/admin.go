package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg/impl"
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/repository"
	"github.com/renfy96/yy-layout-v1/internal/repository/model"
	"github.com/renfy96/yy-layout-v1/internal/repository/query"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	q *query.Query
}

func NewAdminRepository(db *gorm.DB) repository.AdminRepository {
	return repositoryImpl{
		q: query.Use(db),
	}
}

func (r repositoryImpl) GetAdminByAccount(ctx *gin.Context, name string) agg.Admin {
	admin := r.q.Admin
	adminMol, err := admin.WithContext(ctx).Where(admin.Account.Eq(name)).First()
	if err != nil {
		fmt.Println(err)
	}
	return poToAgg(adminMol)
}

func (r repositoryImpl) GetAdminById(ctx *gin.Context, id uint) agg.Admin {
	admin := r.q.Admin
	adminMol, _ := admin.WithContext(ctx).Where(admin.ID.Eq(id)).First()
	return poToAgg(adminMol)
}

func (r repositoryImpl) Save(ctx *gin.Context, admin agg.Admin) (err error) {
	a := r.q.Admin
	err = a.WithContext(ctx).Save(admin.Value())
	return
}

func poToAgg(admin *model.Admin) agg.Admin {
	return impl.NewAdminImpl(admin)
}
