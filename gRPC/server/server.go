package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
	"time"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil { //내장 패키지 사용 및 서버 열기
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		reflection.Register(server)

		go func() { //background 에서만 서버를 돌려야 하기 때문에 thread 생성 (golang의 장점 -> thread 생성 원활)
			log.Println("Start GRPC Server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()
	}

	return nil
}

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token

	s.tokenVerifyMap[token] = data
	return &auth.CreateTokenRes{Auth: data}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token
	res := &auth.VerifyTokenRes{V: &auth.Verify{
		Auth: nil,
	}}

	if authData, ok := s.tokenVerifyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED //관리되지 않는 토큰
	} else if err := s.pasetoMaker.VerifyToken(token); err != nil {
		res.V.Status = auth.ResponseType_FAILED
	} else if authData.ExpireDate < time.Now().Unix() {
		delete(s.tokenVerifyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE //토큰 만료
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
	}

	return res, nil
}
