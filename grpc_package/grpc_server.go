package grpc_package 



//  import folders
import (
	"google.golang.org/grpc"
	"net"
	"fmt"
)

type GRPC_HANDLER interface {
	StartGrpc(port string)(*grpc.Server, net.Listener)
}

type Handle struct {}

func Init () GRPC_HANDLER{return &Handle{}}

func (handlebar *Handle) StartGrpc(port string) (*grpc.Server, net.Listener){

	listen ,err := net.Listen("tcp", ":"+port); if err != nil {
		fmt.Println("Port Allocation:", err.Error())
		return &grpc.Server{}, listen
	}	
	fmt.Println("Listen@:",port)
	return grpc.NewServer(), listen
}

