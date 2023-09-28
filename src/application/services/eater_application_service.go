package services

import (
	"context"
	"errors"
	"fmt"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/eater"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	eatersvc "github.com/DevShuxat/eater-service/src/domain/eater/services"
	"github.com/DevShuxat/eater-service/src/infrastructure/jwt"
	"github.com/DevShuxat/eater-service/src/infrastructure/validator"
)

type EaterAppService interface {
	SignupEater(ctx context.Context, req *pb.SignupEaterRequest) (*pb.SignupEaterResponse, error)
	ConfirmSmsCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc eatersvc.EaterService
	tokenSvc jwt.Service
}

func NewEaterAppService(
	eaterSvc eatersvc.EaterService, 
	tokenSvc jwt.Service,
	) EaterAppService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context,req *pb.SignupEaterRequest) (*pb.SignupEaterResponse,error){
	
	if !validator.ValidateUzPhoneNumber(req.PhoneNumber) {
		return nil,errors.New("Invalid phone number. MUst be of 998{code}{number} format")
	}
	eaterID, err := s.eaterSvc.SignupEater(ctx,req.PhoneNumber)
	
	if err != nil {
		return nil,err
	}
	return &pb.SignupEaterResponse{
		EaterId: eaterID,
	},nil
}


func (s *eaterAppSvcImpl) ConfirmSmsCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error) {

	if req.EaterId == "" {
		return nil,errors.New("invait or missing eater id")
	}

	if req.SmsCode == "" {
		return nil,errors.New("invait or missing sms code")
	}

	profile, err := s.eaterSvc.ConfirmSMSCode(ctx,req.EaterId,req.SmsCode)

	if err != nil {
		return nil,err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.EaterID)

	if err != nil {
		return nil,err
	}

	return &pb.ConfirmSmsCodeResponse{
		Profile: dtos.ToEaterProfilePB(profile),
		Token:   token,
	},nil
 
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context,req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("Invalid or missing profile name: %s",&req.Name)
	}
	profile,err := s.eaterSvc.UpdateEaterProfile(ctx, req.EaterId, req.Name ,req.ImageUrl)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateEaterProfileResponse{
		Profile: dtos.ToEaterProfilePB(profile),
	},nil		
}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context,req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error) {
	if req.EaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater id")
	}
	profile,err := s.eaterSvc.GetEaterProfile(ctx, req.EaterId)
	if err != nil {
		return nil, err
	}

	return &pb.GetEaterProfileResponse{
		Profile: dtos.ToEaterProfilePB(profile),
	},nil		
}