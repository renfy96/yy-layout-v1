package main

import (
	"github.com/renfy96/yy-layout-v1/internal/repository"
	"github.com/renfy96/yy-layout-v1/internal/repository/model"
	"github.com/renfy96/yy-layout-v1/pkg/config"
	"gorm.io/gen"
)

func main() {
	genModel()
}

func genModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb := repository.NewDb(config.NewConfig())
	g.UseDB(gormdb) // reuse your gorm db
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	_ = gormdb.AutoMigrate(
		model.Admin{},
	)
	g.ApplyBasic(
		model.Admin{},
	)
	g.Execute()
}
