package models

import "time"

type Gym struct {
	Id      int    `db:"id"`
	Address string `db:"address"`
	SubLvl  int    `db:"sub_lvl"`
}

type Schedules struct {
	Id        int       `db:"id"`
	GymId     int       `db:"gym_id"`
	DayOfWeek string    `db:"day_of_week"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
