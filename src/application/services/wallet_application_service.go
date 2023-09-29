package services

import (
	"context"
	"errors"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/wallet"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	paymentService "github.com/DevShuxat/eater-service/src/domain/wallet/services"
)

type PaymentApplicationService interface {
	AddCard(ctx context.Context, req *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error)
	GetCard(ctx context.Context, req *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error)
	DeleteCard(ctx context.Context, req *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error)
	ListCard(ctx context.Context, req *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error)
}

type paymentSvc struct {
	paymentsvc paymentService.WalletService
}

func NewPaymentApplicationService(
	paymentsvc paymentService.WalletService,
) PaymentApplicationService {
	return &paymentSvc{
		paymentsvc: paymentsvc,
	}
}

func (s *paymentSvc) AddCard(ctx context.Context, req *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error) {
	if req.CardNumber == "" {
		return nil, errors.New("invalid card number")
	}
	if req.EaterId == "" {
		return nil, errors.New("invalid eater id")
	}

	payment, err := s.paymentsvc.AddCard(ctx, req.CardToken, req.EaterId, req.CardNumber)
	if err != nil {
		return nil, err
	}

	return &pb.AddPaymentCardResponse{
		Card: dtos.ToPaymentPB(payment),
	}, nil

}

func (s *paymentSvc) GetCard(ctx context.Context, req *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error) {
	if req.CardId == "" {
		return nil, errors.New("invalid card id")
	}
	
	payment, err := s.paymentsvc.GetCard(ctx, req.CardId )
	if err != nil {
		return nil, err
	}

	return &pb.GetPaymentCardResponse{
		Card: dtos.ToPaymentPB(payment),
	},nil
}

func (s *paymentSvc) DeleteCard(ctx context.Context, req *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error)

func (s *paymentSvc) ListCard(ctx context.Context, req *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error)
