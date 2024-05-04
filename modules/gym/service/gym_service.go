package service

import (
	"context"
	"github.com/b3liv3r/gym-service/modules/gym/models"
	"github.com/b3liv3r/gym-service/modules/gym/repository"
	"go.uber.org/zap"
)

type GymService struct {
	repo repository.GymerRepository
	log  *zap.Logger
}

func NewGymService(repo repository.GymerRepository, log *zap.Logger) Gymer {
	return &GymService{
		repo: repo,
		log:  log,
	}
}

func (s *GymService) ListGyms(ctx context.Context) ([]models.Gym, error) {
	return s.repo.ListGyms(ctx)
}

func (s *GymService) GetSchedulesForGym(ctx context.Context, gymID int) ([]models.Schedules, error) {
	return s.repo.GetSchedulesForGym(ctx, gymID)
}
