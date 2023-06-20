package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	adminRepo "github.com/renfy96/yy-layout-v1/internal/boudary_context/admin/infrastructure/repository"
	repo "github.com/renfy96/yy-layout-v1/internal/repository"
	"github.com/renfy96/yy-layout-v1/internal/router"
	"github.com/renfy96/yy-layout-v1/pkg/log"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(router.NewServerHTTP)

// RepositorySet 仓储
var RepositorySet = wire.NewSet(
	repo.NewDb,
	adminRepo.NewAdminRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
	))
}
