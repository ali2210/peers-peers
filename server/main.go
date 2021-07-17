package main

import (
	."github.com/ali2210/peers-peers/grpc_package"
	"fmt"
	."github.com/ali2210/peers-peers/protos"
	grpc "google.golang.org/grpc"
	"github.com/ali2210/peers-peers/contacts"
)

var  (
	grpcHandler Handle = Handle{}
	friend contacts.Friends
)


func main() {
	
	port := "9000"
	fmt.Println("GRPC started...")
	
	
	server, listen := grpcHandler.StartGrpc(port)
	 contacts.RegisterPhoneService(grpc.NewServer(), &Phonebook_Server{})
	 err := server.Serve(listen); if err != nil{
		fmt.Println("Port Busy:", err.Error())
		return
	}
	
}