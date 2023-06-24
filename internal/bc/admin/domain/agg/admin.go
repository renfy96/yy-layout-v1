package agg

import "github.com/renfy96/yy-layout-v1/internal/repository/model"

type Admin interface {
	Value() *model.Admin
	IsExist() bool
	CheckPassword(pwd string) bool
	UpdatePassword(password string)
}
