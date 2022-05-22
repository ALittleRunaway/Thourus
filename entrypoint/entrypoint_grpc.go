package entrypoint

import (
	"context"
	"github.com/spf13/viper"
	"thourus-api/domain/entity"
	"thourus-api/domain/usecase"
	"thourus-api/gateway"
	pb "thourus-api/infrastructure/grpc/proto"
)

type GrpcServer struct {
	pb.UnimplementedCryptoCoreServer
}

func (s *GrpcServer) Mine(ctx context.Context, request *pb.MineRequest) (response *pb.MineResponse, err error) {

	receivedMessage := entity.MineRequestMessage{
		Bytes: request.Bytes,
		Rule:  request.Rule,
	}

	cryptoGw := gateway.NewCryptoGateway(viper.GetString("cryptocore.secret_string"))
	uc := usecase.NewCryptoUseCase(cryptoGw)

	proofOfWork, err := uc.Mine(receivedMessage.Bytes, receivedMessage.Rule)
	if err != nil {
		return nil, err
	}

	response = &pb.MineResponse{
		Pow: proofOfWork,
	}

	return response, nil
}

func (s *GrpcServer) Validate(ctx context.Context, request *pb.ValidateRequest) (response *pb.ValidateResponse, err error) {

	receivedMessage := entity.ValidateRequestMessage{
		Bytes: request.Bytes,
		Hash:  request.Hash,
		PoW:   request.Pow,
	}

	cryptoGw := gateway.NewCryptoGateway(viper.GetString("cryptocore.secret_string"))
	uc := usecase.NewCryptoUseCase(cryptoGw)

	valid, err := uc.Validate(receivedMessage.Bytes, receivedMessage.Hash, receivedMessage.PoW)
	if err != nil {
		return nil, err
	}

	response = &pb.ValidateResponse{
		Valid: valid,
	}

	return response, nil
}
