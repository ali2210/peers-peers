package main

import (
	"context"
	"time"
	pb "github.com/ali2210/pal/address"
	"fmt"
)

const(
	Address string = "localhost:5000"
)

func main() {
	
	connection , err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return 
	}
	defer connection.Close()

	client := pb.NewContractClient(connection)
	ctx , cancel := contetxt.WithTimeout(contetxt.Background, time.Seconds)
	defer cancel
	contact , err := client.AddPeers(ctx, &pb.Peers{
		Name : "Ali",
		Eve : false,
		Pwallet : "0xBFf88f5E697852246222E85BF0DACa6af0FA48Cf"
	})
	fmt.Println("CONTACT:", contact)
	if err != nil {
		return 
	}
	response ,err := client.DisplayPeers(ctx, contact)
	if err != nil {
		return 
	}
	fmt.Println("Display info", response)
	deleteRes , err := client.DeletePeers(ctx, &pb.Peers{
		Name : "Ali",
		Eve : false,
		Pwallet : "0xBFf88f5E697852246222E85BF0DACa6af0FA48Cf"
	})
	if err != nil {
		return 
	}
	fmt.Println("Delete info", deleteRes)
	
}