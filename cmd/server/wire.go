//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	adminSvc "github.com/renfy96/yy-layout-v1/internal/application/admin/service"
	adminApp "github.com/renfy96/yy-layout-v1/internal/bc/admin/application"
	adminRepo "github.com/renfy96/yy-layout-v1/internal/bc/admin/infrastructure/repository"
	repo "github.com/renfy96/yy-layout-v1/internal/repository"
	"github.com/renfy96/yy-layout-v1/internal/router"
	"github.com/renfy96/yy-layout-v1/pkg/log"
	"github.com/spf13/viper"
)

// RepositorySet 仓储
var RepositorySet = wire.NewSet(
	repo.NewDb,
	adminRepo.NewAdminRepository,
)

// 请求服务
var repoSvc = wire.NewSet(
	adminApp.NewCommandService,
	adminApp.NewQueryService,
)

// SvcSet 业务服务
var SvcSet = wire.NewSet(
	adminSvc.NewService,
)

var ServerSet = wire.NewSet(router.NewServerHTTP)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		RepositorySet,
		repoSvc,
		SvcSet,
		ServerSet,
	))
}
