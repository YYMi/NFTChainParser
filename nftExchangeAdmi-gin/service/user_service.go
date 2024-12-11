package service

import (
	errors2 "errors"
	"nftExchangeAdmi-gin/errors"
	"nftExchangeAdmi-gin/types"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetAllProducts 处理具体业务逻辑
func (s *UserService) GetAllProducts(name *string, aget *int) (*types.UserVO, *errors.MyError) {
	if *name == "" || *aget <= 0 {
		return nil, errors.SysError(errors2.New("invalid user data"))
	}
	return &types.UserVO{
		ID:      1,
		Name:    *name,
		Message: "Successfully fetched product list",
	}, nil
}
