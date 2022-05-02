package main

import (
	"context"
	"encoding/json"
	"fmt"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net"
	pb "simple-grpc-trial/student"
	"sync"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu sync.Mutex // antisipasi race condition
	students []*pb.Student // for JSON results
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error){
	fmt.Println("Incoming request")
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}

	return nil, status.Error(codes.NotFound, "Student email unknown")
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/students.json")
	if err != nil{
		log.Fatal("Error in read file.\n", err.Error())
	}

	if err:= json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error in decoding data.\n", err.Error())
	}

}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main(){
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in listen.\n", err.Error())
	}

	grpcServer := grpc2.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	// run server
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Error when serve grpc.\n", err.Error())
	}
}
