package service

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/boudary_context/admin/application"
)

type Service struct {
	command *application.CommandService
}

func NewService(svc *application.CommandService) *Service {
	return &Service{command: svc}
}

func (svc Service) Login(ctx *gin.Context) {

}