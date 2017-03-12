package main

import (
	"encoding/json"
	"log"

	userbase "hellomall/userbase/srv/proto"
    mallhome "hellomall/mallhome/srv/proto"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"

    "github.com/micro/go-grpc"
)

type BaseData struct {}




func main() {
	service := grpc.NewService(
        micro.Name("go.micro.srv.mallhome"),
        micro.RegisterTTL(time.Second*30),
        micro.RegisterInterval(time.Second*10),
    )

    // optionally setup command line usage
    service.Init()

    // Register Handlers
    mallhome.RegisterMallHomeHandler(service.Server(), new(BaseData))

    // Run server
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
