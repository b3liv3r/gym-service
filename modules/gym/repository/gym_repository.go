package repository

import (
	"context"
	"github.com/b3liv3r/gym-service/modules/gym/models"
	"github.com/jmoiron/sqlx"
)

type GymRepositoryDB struct {
	db *sqlx.DB
}

func NewGymRepositoryDB(db *sqlx.DB) GymerRepository {
	return &GymRepositoryDB{db: db}
}

func (r *GymRepositoryDB) ListGyms(ctx context.Context) ([]models.Gym, error) {
	var gyms []models.Gym
	err := r.db.SelectContext(ctx, &gyms, "SELECT id, address, sub_lvl FROM gyms")
	if err != nil {
		return nil, err
	}
	return gyms, nil
}

func (r *GymRepositoryDB) GetSchedulesForGym(ctx context.Context, gymID int) ([]models.Schedules, error) {
	var schedules []models.Schedules
	err := r.db.SelectContext(ctx, &schedules, "SELECT id, gym_id, day_of_week, start_time, end_time FROM schedules WHERE gym_id = $1", gymID)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
