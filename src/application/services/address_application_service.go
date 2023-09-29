package services

import (
	"context"
	"errors"
	"fmt"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/address"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	addresssvc "github.com/DevShuxat/eater-service/src/domain/address/services"
)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error)
	UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error)
	GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error)
	ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListAddressResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addresssvc.AddressService
}

func NewAddressApplicationService(
	addressSvc addresssvc.AddressService,
) AddressApplicationService {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {

	if req.EaterId == "" {
		return nil, errors.New("invalid or missing eater id")
	}

	if req.Name == "" {
		return nil, errors.New("invalid or missing address name")
	}

	if req.Longitude == 0 {
		return nil, errors.New("invalid or missing longitude")
	}

	if req.Latitude == 0 {
		return nil, errors.New("invalid or missing latitude")
	}

	address, err := s.addressSvc.SaveAddress(ctx, req.Name, req.EaterId, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return &pb.AddAddressResponse{
		Address: dtos.ToAddressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {

	if req.AddressId == "" {
		return nil, errors.New("invalid or missing address id")
	}

	if req.Name == "" {
		return nil, errors.New("invalid or missing address name")
	}

	if req.Longitude == 0 {
		return nil, errors.New("invalid or missing longitude")
	}

	if req.Latitude == 0 {
		return nil, errors.New("invalid or missing latitude")
	}

	address, err := s.addressSvc.UpdateAddress(ctx, req.AddressId, req.Name, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressResponse{
		Address: dtos.ToAddressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {

	if req.GetAddressId() == "" {
		return nil, errors.New("invalid or missing address id")
	}

	err := s.addressSvc.DeleteAddress(ctx, req.GetAddressId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAddressResponse{}, nil
}

func (s *addressAppSvcImpl) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {

	if req.GetAddressId() == "" {
		return nil, fmt.Errorf("Invalid or missing addressId: %s", req.GetAddressId())
	}

	address, err := s.addressSvc.GetAddress(ctx, req.GetAddressId())
	if err != nil {
		return nil, err
	}

	return &pb.GetAddressResponse{
		Address: dtos.ToAddressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListAddressResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing addeaterIDressId: %s", eaterID)
	}

	addresses, err := s.addressSvc.ListAddressByEater(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return dtos.NewListAddressResponse(addresses), nil
}
