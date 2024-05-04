package service

import (
	"context"
	"github.com/b3liv3r/gym-service/modules/gym/models"
)

type Gymer interface {
	ListGyms(ctx context.Context) ([]models.Gym, error)
	GetSchedulesForGym(ctx context.Context, gymId int) ([]models.Schedules, error)
}
