package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
)

type GRPCServer struct {
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil { //내장 패키지 사용 및 서버 열기
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)

		////AuthServiceServer
		//auth.RegisterAuthServiceServer(server, &GRPCServer{
		//		pasetoMaker: paseto.NewPasetoMaker(),
		//		tokenVerifyMap: make(map[string]*auth.AuthData),
		//}

		reflection.Register(server)

		go func() { //background 에서만 서버를 돌려야 하기 때문에 thread 생성 (golang의 장점 -> thread 생성 원활)
			log.Println("Start GRPC Server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()

		//RegisterAuthServiceServer
	}

	return nil
}
