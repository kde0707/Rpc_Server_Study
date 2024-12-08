package main

import (
	"flag"
	"rpc-server/cmd"
	"rpc-server/config"
	"rpc-server/gRPC/server"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "config path") // configFlag 를 통해 path 를 유동적으로 지정할 수 있다.

func main() {
	flag.Parse()
	//fmt.Println(*configFlag) // go run . -config = 지정경로 => aws, instance 로 바꿨을 때 경로 변경이 용이하다.
	cfg := config.NewConfig(*configFlag)

	if err := server.NewGRPCServer(cfg); err != nil {
		panic(err)
	} else {
		time.Sleep(1e9)
		cmd.NewApp(cfg)
	}
}
