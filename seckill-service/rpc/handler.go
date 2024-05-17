package rpc

import (
	"context"
	itemcenter "go_code/seckill-service/rpc/kitex_gen/itemcenter"
	"go_code/seckill-service/service"
)

// ItemcenterImpl implements the last service interface defined in the IDL.
type ItemcenterImpl struct{}

// Register implements the ItemcenterImpl interface.
func (s *ItemcenterImpl) Register(ctx context.Context, u *itemcenter.User) (err error) {
	// TODO: Your code here...
	return service.Register(u)
}

// Login implements the ItemcenterImpl interface.
func (s *ItemcenterImpl) Login(ctx context.Context, u *itemcenter.User) (resp *itemcenter.User, err error) {
	// TODO: Your code here...
	return service.Login(u)
}

// ParseToken implements the ItemcenterImpl interface.
func (s *ItemcenterImpl) ParseToken(ctx context.Context, token string, typ int64) (resp string, err error) {
	// TODO: Your code here...
	return service.ParseToken(token, typ)
}

// CreateToken implements the ItemcenterImpl interface.
func (s *ItemcenterImpl) CreateToken(ctx context.Context, u *itemcenter.User, typ int64) (resp string, err error) {
	// TODO: Your code here...
	return service.CreateToken(u, typ)
}

// SecKill implements the ItemcenterImpl interface.
func (s *ItemcenterImpl) SecKill(ctx context.Context, u *itemcenter.User) (resp bool, err error) {
	// TODO: Your code here...
	return service.SecKill(u)
}
