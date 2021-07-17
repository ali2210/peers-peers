package main

import (
	."github.com/ali2210/peers-peers/grpc_package"
	"fmt"
	."github.com/ali2210/peers-peers/protos"
	"context"
	"github.com/ali2210/peers-peers/contacts"
)

var  (
	grpcHandler Handle = Handle{}
	friend contacts.Friends
)

func main() {
	port := "8900"
	fmt.Println("GRPC started...")
	server, listen := grpcHandler.StartGrpc(port)
	dairy := Phonebook_Server{}
	fmt.Println("PhoneBook:", dairy)
	profile := dairy.AddContact(context.Background(), contacts.Friends{Id: 0, Name: "Ali", Email:"alideveloper95@gmail.com", Phone:"+92-",})
	fmt.Println("update:", profile.GetPalx())
	err := server.Serve(listen); if err != nil{
		fmt.Println("Port Busy:", err.Error())
		return
	}
	
}