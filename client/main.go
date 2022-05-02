package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"simple-grpc-trial/student"
	"time"
)

func getDataStudentByEmail(client student.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := student.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatal("error when getting student by email.\n", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatal("Error in connection dial.\n", err.Error())
	}

	defer conn.Close()

	client := student.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "saji@gmail.co")
	getDataStudentByEmail(client, "asasas@gmail.co")

}
