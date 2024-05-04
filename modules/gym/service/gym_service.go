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
	gyms, err := s.repo.ListGyms(ctx)
	if err != nil {
		s.log.Error("Error listing gyms", zap.Error(err))
	}
	return gyms, err
}

func (s *GymService) GetSchedulesForGym(ctx context.Context, gymID int) ([]models.Schedules, error) {
	schedules, err := s.repo.GetSchedulesForGym(ctx, gymID)
	if err != nil {
		s.log.Error("Error getting gym schedules", zap.Error(err))
		return nil, err
	}
	return schedules, err
}
