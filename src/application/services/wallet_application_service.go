package services
import (
	"context"
	dtos "github.com/DevShuxat/eater-service/src/application/dtos/wallet"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	paymentService "github.com/DevShuxat/eater-service/src/domain/wallet/service"
)

type PaymentService interface {
	AddCard(ctx context.Context, req *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error)
	GetCard(ctx context.Context, req *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error)
	DeleteCard(ctx context.Context, req *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error)
	ListCard(ctx context.Context, req *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error)
}

type paymentSvc struct {
	ratingSvc PaymentService.PaymentService
}

func NewPaymentApplicationService(
	ratingSvc PaymentService.paymentService,
) PaymentApplicationService {
	return &paymentSvc{
		paymentSvc: paymentSvc,
	}
}

