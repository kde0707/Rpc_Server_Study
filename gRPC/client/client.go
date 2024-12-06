package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
	"time"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	if client, err := grpc.Dial(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil { //WithTransportCredentials 보안 연결을 사용하지 않는 경우
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(c.client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}

	//auth.NewAuthServiceClient()
	return c, nil
}

func (g *GRPCClient) CreateAuth(name string) (*auth.AuthData, error) {
	now := time.Now()
	expiredTime := now.Add(30 * time.Minute)

	a := &auth.AuthData{
		Name:       name,
		CreateDate: now.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	if token, err := g.pasetoMaker.CreateNewToken(a); err != nil {
		return nil, err
	} else {
		a.Token = token

		if res, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: a}); err != nil {
			return nil, err
		} else {
			return res.Auth, nil
		}
	}
	return nil, nil
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {
	if res, err := g.authClient.VerifyToken(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res.V, nil
	}
	return nil, nil
}
