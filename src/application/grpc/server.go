package grpc

import (
	"context"

	"github.com/DevShuxat/eater-service/src/application/services"
)
type Server struct{
	pb.EaterServiceServer
	eaterApp services.EaterApplicationService
	addressApp services.AddressAppService
	deliveryRatingApp services.DeliveryRatingAppService
	restaurantRatingApp services.RestaurantRatingAppService
}

func NewServer(
	eaterApp services.EaterApplicationService,
	addressApp services.AddressAppService,
	deliveryRatingApp services.DeliveryRatingAppService,
	restaurantRatingApp services.RestaurantRatingAppService,
) *Server{
	return &Server{
		eaterApp: eaterApp,
		addressApp: addressApp,
		deliveryRatingApp:deliveryRatingApp,
		restaurantRatingApp:restaurantRatingApp,
	}
}

func (s *Server) SignupEater(ctx context.Context, r *pb.SignupEaterRequest)(*pb.SignupEaterResponse,error){
	return s.eaterApp.SignupEater(ctx,r)
}

func (s *Server) ConfirmSmsCode(ctx context.Context, r *pb.ConfirmSmsCodeRequest)(*pb.ConfirmSmsCodeResponse,error){
	return s.eaterApp.ConfirmSMSCode(ctx,r)
}

func (s *Server) UpdateEaterProfile(ctx context.Context, r *pb.UpdateEaterProfileRequest)(*pb.UpdateEaterProfileResponse,error){
	return s.eaterApp.UpdateEaterProfile(ctx,r)
}

func (s *Server) GetEaterProfile(ctx context.Context, r *pb.GetEaterProfileRequest)(*pb.GetEaterProfileResponse,error){
	return s.eaterApp.GetEaterProfile(ctx,r)
}

func (s *Server) SaveAddress(ctx context.Context, r *pb.AddAddressRequest)(*pb.AddAddressResponse, error){
	return s.addressApp.SaveAddress(ctx,r)
}

func (s *Server) Update(ctx context.Context, r *pb.UpdateAddressRequest)(*pb.UpdateAddressResponse, error){
	return s.addressApp.UpdateAddress(ctx,r)
}

func (s *Server) GetAddress(ctx context.Context, r *pb.GetAddressRequest)(*pb.GetAddressResponse, error){
	return s.addressApp.GetAddress(ctx,r)
}

func (s *Server) DeleteAddress(ctx context.Context, r *pb.DeleteAddressRequest)(*pb.DeleteAddressResponse, error){
	return s.addressApp.DeleteAddress(ctx,r)
}
// ----------------------------------------------------------
// start Deliver Rating 
func (s *Server) RateDelivery(ctx context.Context, r *pb.RateDeliveryRequest)(*pb.RateDeliveryResponse,error){
	return s.deliveryRatingApp.SaveRating(ctx,r)
}

func (s *Server) UpdateDeliveryRating(ctx context.Context, r *pb.UpdateDeliveryRatingRequest)(*pb.UpdateDeliveryRatingResponse,error){
	return s.deliveryRatingApp.UpdateRating(ctx,r)
}

func (s *Server) ListDeliveryRatingByEater(ctx context.Context, r *pb.ListDeliveryRatingByEaterRequest)(*pb.ListDeliveryRatingByEaterResponse,error){
	return s.deliveryRatingApp.GetRatingsByEater(ctx,r)
}

func (s *Server) GetDeliveryRatingByOrder(ctx context.Context, r *pb.GetDeliveryRatingByOrderRequest)(*pb.GetDeliveryRatingByOrderResponse,error){
	return s.deliveryRatingApp.GetRatingByOrder(ctx,r)
}

// end Delivery Rating
// -----------------------------------------------------------


// start Restaurant rating
// -----------------------------------------------------------
func (s *Server) RateRestaurant(ctx context.Context, r *pb.RateRestaurantRequest)(*pb.RateRestaurantResponse,error){
	return s.restaurantRatingApp.SaveRating(ctx,r)
}

func (s *Server) UpdateRestaurantRating(ctx context.Context, r *pb.UpdateRestaurantRatingRequest)(*pb.UpdateRestaurantRatingResponse,error){
	return s.restaurantRatingApp.UpdateRating(ctx,r)
}

func (s *Server) ListRestaurantRatingByEater(ctx context.Context, r *pb.ListRestaurantRatingByEaterRequest)(*pb.ListRestaurantRatingByEaterResponse,error){
	return s.restaurantRatingApp.GetRatingByEater(ctx,r)
}
// end Restaurant rating
// -----------------------------------------------------------