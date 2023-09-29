package wallet

import (
	"time"

	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
)

func ToPaymentPB(payment *models.PaymentCard) *pb.PaymentCardView {
	return &pb.PaymentCardView{
		Id:         payment.ID,
		EaterId:    payment.EaterID,
		Number:     payment.Number,
		IsVerified: payment.IsVerified,
		CreatedAt:  payment.CreatedAt.Format(time.RFC3339),
	}
}
