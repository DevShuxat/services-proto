package dtos

import (
	"time"

	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	"github.com/DevShuxat/eater-service/src/domain/eater/models"
)
func ToEaterProfilePB(profile *models.EaterProfile) *pb.EaterProfile{
	return &pb.EaterProfile{
		EaterId: profile.EaterID,
		PhoneNumber: profile.PhoneNumber,
		Name: profile.Name,
		IsPhoneNumberConfirmed: profile.IsPhoneNumberConfirmed,
		CreatedAt: profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt: profile.UpdatedAt.Format(time.RFC3339),
	}
}