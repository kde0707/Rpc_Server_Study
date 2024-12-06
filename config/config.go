package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Paseto struct {
		Key string
	}
}

func NewConfig(path string) *Config {
	c := new(Config) //config 경로 받아오기

	if file, err := os.Open(path); err != nil { //os를 통해 toml 파일 열기
		panic(err) //실행중이던 것을 모두 멈추고, return
	} else {
		defer file.Close() //defer(지연실행) : 모든 함수를 다 호출하고 실행하는 것

		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			defer file.Close()

			if err = toml.NewDecoder(file).Decode(c); err != nil {
				panic(err)
			} else {
				return c
			}
		}
	}
}
