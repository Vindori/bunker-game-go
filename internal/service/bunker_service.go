package service

import (
	"context"
	pb "github.com/Vindori/bunker-game-go/internal/pb/api"
)

type BunkerService struct {
	pb.UnimplementedBunkerGameServiceServer
}

// TODO: Наебашить интерфейс для взаимодействия с хранилищем + реализовать рерпу для хранения игроков

func (s *BunkerService) RevealPlayerCharacteristic(ctx context.Context, request *pb.RevealPlayerCharacteristicRequest) (*pb.RevealPlayerCharacteristicResponse, error) {
	resp := &pb.RevealPlayerCharacteristicResponse{
		Characteristic: &pb.PlayerCharacteristic{
			Value: &pb.PlayerCharacteristic_Orientation_{Orientation: pb.PlayerCharacteristic_ORIENTATION_HOMO},
		},
	}
	return resp, nil
}

func NewBunkerService() *BunkerService {
	return &BunkerService{}
}
