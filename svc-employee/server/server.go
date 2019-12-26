package server

import (
	"context"
	"github.com/TheShifter/grpcRest/data"
	employee "github.com/TheShifter/grpcRest/svc-employee/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type emplServer struct {}

func (e *emplServer) List(ctx context.Context, req *employee.GetEmployeesRequest) (*employee.GetEmployeesResponse, error) {
	dataCon := data.ConnectToDb()
	empls, err := dataCon.GetEmployeeList(req.Id)
	if err != nil {
		log.Fatalln(err)
	}
	var response []*employee.Employee
	for i := 0; i < len(empls); i++ {
		empl := &employee.Employee{
			Id: empls[i].Id,
			Name: empls[i].Name,
			Age: empls[i].Age,
			UnitId: empls[i].Unit_id,
		}
		response = append(response, empl)
	}
	return &employee.GetEmployeesResponse{Emp:response}, nil
}

func (e *emplServer) Read(ctx context.Context, req *employee.GetOneEmployeeRequest) (*employee.GetOneEmployeeResponse, error) {
	dataCon := data.ConnectToDb()
	empl, err := dataCon.GetEmployeeById(req.Id, req.Id1)
	if err != nil {
		log.Fatalln(err)
	}
	response:= &employee.Employee{
		Id: empl.Id,
		Name: empl.Name,
		Age: empl.Age,
		UnitId: empl.Unit_id,
	}
	return &employee.GetOneEmployeeResponse{Emp: response}, nil
}

func (e *emplServer) Create(ctx context.Context, req *employee.CreateEmployeeRequest) (*employee.CreateEmployeeResponse, error) {
	dataCon := data.ConnectToDb()
	details := &data.Employee{Id:req.Emp.Id, Name: req.Emp.Name, Age: req.Emp.Age, Unit_id: req.Emp.UnitId}
	created := dataCon.CreateEmployee(req.Id, details)
	response := &employee.Employee{Id:created.Id, Name: created.Name, Age: created.Age, UnitId: created.Unit_id}
	return &employee.CreateEmployeeResponse{Emp: response}, nil
}

func (e *emplServer) Update(ctx context.Context, req *employee.UpdateEmployeeRequest) (*employee.UpdateEmployeeResponse, error) {
	dataCon := data.ConnectToDb()
	details := &data.Employee{Id:req.Emp.Id, Name: req.Emp.Name, Age: req.Emp.Age, Unit_id: req.Emp.UnitId}
	updated := dataCon.UpdateEmployee(req.Id, req.Id1, details)
	response := &employee.Employee{Id: updated.Id, Name: updated.Name, Age: updated.Age, UnitId: updated.Unit_id}
	return &employee.UpdateEmployeeResponse{Emp: response}, nil
}

func (e *emplServer) Delete(ctx context.Context, req *employee.DeleteEmployeeRequest) (*employee.DeleteEmployeeResponse, error) {
	dataCon := data.ConnectToDb()
	err := dataCon.DeleteEmployee(req.Id, req.Id1)
	if err != nil {
		log.Fatalln(err)
	}
	return &employee.DeleteEmployeeResponse{}, nil
}

func NewService(serviceAddress string) {
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	employee.RegisterEmployeeServiceServer(server, &emplServer{})
	reflection.Register(server)
	server.Serve(listener)
}
