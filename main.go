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
	fmt.Println("Data serialization started with protobuf ...")
	dairy := Phonebook_Server{}
	profile := dairy.AddContact(context.Background(), &contacts.Friends{Id: 0, Name: "Ali", Email:"alideveloper95@gmail.com", Phone:"+92-",})
	fmt.Println("update:", profile.GetPalx())
	
}