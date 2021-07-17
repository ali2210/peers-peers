package main

import (
	"time"
	"github.com/ali2210/peers-peers/contacts"
	grpc "google.golang.org/grpc"
	."github.com/ali2210/peers-peers/protos"
	"context"
	"fmt"
)
var(
	address string = "localhost:9000"
)
func main() {
	
	conn, err := grpc.Dial(address, grpc.WithInsecure(),grpc.WithBlock()); if err != nil {
		fmt.Println("Connection :", conn)
		return 
	}
	defer conn.Close()
	
	ctx , cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	client_grpc := contacts.NewClientServiceInit(conn)
	dairy := Phonebook_Server{}
	profile := dairy.AddContact(context.Background(), &contacts.Friends{Id: 0, Name: "Ali", Email:"alideveloper95@gmail.com", Phone:"+92-",})
	user := client_grpc.AddContact(ctx, &contacts.Friends{Id: 0, Name: "Ali", Email:"alideveloper95@gmail.com", Phone:"+92-",})
	fmt.Println("proto_user_response:", profile)
	fmt.Println("Response:", user.GetPalx())
}