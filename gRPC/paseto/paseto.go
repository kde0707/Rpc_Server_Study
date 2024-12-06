package paseto

import (
	"github.com/o1egl/paseto"
	"rpc-server/config"
	auth "rpc-server/gRPC/proto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

// payload를 추가함으로써 새로운 토큰 제작
func (m *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	return "", nil
}

// 토큰 수령 후 검증
func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
