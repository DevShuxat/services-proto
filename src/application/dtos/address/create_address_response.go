package dtos

import (
	"time"

	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	"github.com/DevShuxat/eater-service/src/domain/address/models"
)

func ToAddressPB(address *models.Address) *pb.Address {
	return &pb.Address{
		Id:      address.ID,
		EaterId: address.EaterID,
		Name:    address.Name,
		Location: &pb.Location{
			Longitude: address.Location.Longitude,
			Latitude:  address.Location.Latitude,
		},
		CreatedAt: address.CreatedAt.Format(time.RFC3339),
		UpdatedAt: address.UpdatedAt.Format(time.RFC3339),
	}
}
