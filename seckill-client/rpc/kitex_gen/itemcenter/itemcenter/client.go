// Code generated by Kitex v0.9.1. DO NOT EDIT.

package itemcenter

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	itemcenter "go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (err error)
	Login(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (r *itemcenter.User, err error)
	ParseToken(ctx context.Context, token string, typ int64, callOptions ...callopt.Option) (r string, err error)
	CreateToken(ctx context.Context, u *itemcenter.User, typ int64, callOptions ...callopt.Option) (r string, err error)
	SecKill(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (r bool, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kItemcenterClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kItemcenterClient struct {
	*kClient
}

func (p *kItemcenterClient) Register(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, u)
}

func (p *kItemcenterClient) Login(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (r *itemcenter.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, u)
}

func (p *kItemcenterClient) ParseToken(ctx context.Context, token string, typ int64, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ParseToken(ctx, token, typ)
}

func (p *kItemcenterClient) CreateToken(ctx context.Context, u *itemcenter.User, typ int64, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateToken(ctx, u, typ)
}

func (p *kItemcenterClient) SecKill(ctx context.Context, u *itemcenter.User, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SecKill(ctx, u)
}