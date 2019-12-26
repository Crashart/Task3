package server

import (
	"context"
	"github.com/TheShifter/grpcRest/data"
	unit "github.com/TheShifter/grpcRest/svc-unit/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type unitServer struct {}

func (u *unitServer) List(ctx context.Context, req *unit.GetUnitsRequest) (resp *unit.GetUnitsResponse, err error) {
	dataCon := data.Connect()
	units, err := dataCon.GetUnitList()
	if err != nil {
		log.Fatalln(err)
	}
	var response []*unit.Unit
	for i := 0; i < len(units); i++ {
		unit := &unit.Unit{
			Id: units[i].Id,
			Name: units[i].Name,
		}
		response = append(response, unit)
	}
	return &unit.GetUnitsResponse{Units: response}, nil
}

func (u *unitServer) Read(ctx context.Context, req *unit.GetOneUnitRequest) (resp *unit.GetOneUnitResponse, err error) {
	dataCon := data.Connect()
	oneUnit, err := dataCon.GetUnitById(req.Id)
	if err != nil {
		log.Fatalln(err)
	}
	response:= &unit.Unit{
		Id: oneUnit.Id,
		Name: oneUnit.Name,
	}
	return &unit.GetOneUnitResponse{Unit:response}, nil
}

func (u *unitServer) Create(ctx context.Context, req *unit.CreateUnitRequest) (resp *unit.CreateUnitResponse, err error) {
	dataCon := data.Connect()
	details := &data.Unit{Id: req.Unit.Id, Name: req.Unit.Name}
	created := dataCon.CreateUnit(details)
	response := &unit.Unit{Id: created.Id, Name: created.Name}
	return &unit.CreateUnitResponse{Unit: response}, nil
}

func (u *unitServer) Update(ctx context.Context, req *unit.UpdateUnitRequest) (resp *unit.UpdateUnitResponse, err error) {
	dataCon := data.Connect()
	details := &data.Unit{Id: req.Unit.Id, Name: req.Unit.Name}
	updated := dataCon.UpdateUnit(req.Id, details)
	response := &unit.Unit{Id: updated.Id, Name: updated.Name}
	return &unit.UpdateUnitResponse{Unit: response}, nil
}

func (u *unitServer) Delete(ctx context.Context, req *unit.DeleteUnitRequest) ( *unit.DeleteUnitResponse,  error) {
	dataCon := data.Connect()
	err := dataCon.DeleteUnit(req.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return &unit.DeleteUnitResponse{}, nil
}

func NewService(serviceAddress string) {
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	unit.RegisterUnitServiceServer(server, &unitServer{})
	reflection.Register(server)
	server.Serve(listener)
}