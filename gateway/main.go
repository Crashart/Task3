package main

import (
	"context"
	"flag"
	employee "github.com/TheShifter/grpcRest/svc-employee/pb"
	eServer "github.com/TheShifter/grpcRest/svc-employee/server"
	unit "github.com/TheShifter/grpcRest/svc-unit/pb"
	uServer "github.com/TheShifter/grpcRest/svc-unit/server"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	unitAddress = "localhost:9090"
	emplAddress = "localhost:9091"
)


func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go uServer.NewService(unitAddress)
	go eServer.NewService(emplAddress)
	mux := runtime.NewServeMux()
	UConn, err := grpc.Dial(unitAddress, grpc.WithInsecure())
	EConn, err := grpc.Dial(emplAddress, grpc.WithInsecure())
	err = unit.RegisterUnitServiceHandler(ctx, mux, UConn)
	err = employee.RegisterEmployeeServiceHandler(ctx, mux, EConn)
	if err != nil {
		log.Fatalln(err)
	}
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

