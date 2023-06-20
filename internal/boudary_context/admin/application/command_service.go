package application

import "github.com/renfy96/yy-layout-v1/internal/boudary_context/admin/domain/repository"

type CommandService struct {
	adminRepo repository.AdminRepository
}

func NewCommandService(repo repository.AdminRepository) *CommandService {
	return &CommandService{adminRepo: repo}
}
