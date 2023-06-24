package impl

import (
	"github.com/renfy96/yy-layout-v1/internal/bc/admin/domain/agg"
	"github.com/renfy96/yy-layout-v1/internal/repository/model"
)

type adminImpl struct {
	value *model.Admin
}

func NewAdminImpl(v *model.Admin) agg.Admin {
	return &adminImpl{
		value: v,
	}
}

func (a adminImpl) CheckPassword(pwd string) bool {
	return a.value.Password == pwd
}

func (a adminImpl) UpdatePassword(password string) {
	//TODO implement me
	panic("implement me")
}

func (a adminImpl) IsExist() bool {
	if a.value == nil {
		return false
	}
	return a.value.ID > 0
}

func (a adminImpl) Value() *model.Admin {
	return a.value
}
