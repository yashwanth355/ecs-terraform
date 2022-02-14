package user

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/userrepo"
)

type Service interface {
	CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error)
	AdminConfirmSignUp(ctx context.Context, resquest model.Credentials) (*model.AdminconfirmSignUpRes, error)
}

type service struct {
	userRepo userrepo.UserRepo
}

func NewuserService() *service {
	return &service{
		userRepo: userrepo.NewUserRepo(),
	}
}

func (s service) CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error) {
	user := entity.User{

		Id:  resquest.Id,
		Id1: resquest.Id1,
		Id2: resquest.Id2,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return &model.UserResponse{Status: 232}, nil
}

func (s *service) AdminConfirmSignUp(ctx context.Context, resquest model.Credentials) (*model.AdminconfirmSignUpRes, error) {
	//written all the logic here

	return nil, nil
}
