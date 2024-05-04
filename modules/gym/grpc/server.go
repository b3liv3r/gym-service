package server

import (
	"context"
	"github.com/b3liv3r/gym-service/modules/gym/service"
	gymv1 "github.com/b3liv3r/protos-for-gym/gen/go/gym"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GymRPCServer struct {
	gymv1.UnimplementedGymServer
	srv service.Gymer
}

func NewGymRPCServer(srv service.Gymer) gymv1.GymServer {
	return &GymRPCServer{srv: srv}
}

func (s *GymRPCServer) List(ctx context.Context, req *gymv1.ListRequest) (*gymv1.ListResponse, error) {
	gyms, err := s.srv.ListGyms(ctx)
	if err != nil {
		return nil, err
	}

	var gymList []*gymv1.Gym
	for _, gym := range gyms {
		gymList = append(gymList, &gymv1.Gym{
			GymId:   int64(gym.Id),
			Address: gym.Address,
			SubLvl:  int64(gym.SubLvl),
		})
	}

	return &gymv1.ListResponse{GymList: gymList}, nil
}

func (s *GymRPCServer) GetSchedules(ctx context.Context, req *gymv1.GetSchedulesRequest) (*gymv1.GetSchedulesResponse, error) {
	schedules, err := s.srv.GetSchedulesForGym(ctx, int(req.GymId))
	if err != nil {
		return nil, err
	}

	var scheduleList []*gymv1.Schedule
	for _, schedule := range schedules {
		startTime := timestamppb.New(schedule.StartTime)
		endTime := timestamppb.New(schedule.EndTime)
		scheduleList = append(scheduleList, &gymv1.Schedule{
			ScheduleId: int64(schedule.Id),
			GymId:      int64(schedule.GymId),
			DayOfWeek:  schedule.DayOfWeek,
			StartTime:  startTime,
			EndTime:    endTime,
		})
	}

	return &gymv1.GetSchedulesResponse{ScheduleList: scheduleList}, nil
}
