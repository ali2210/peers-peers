package grpc_package 



//  import folders
import (
	"google.golang.org/grpc"
)

type GRPC_HANDLER interface {
	StartGrpc()
}

type Handle struct {}

func Init () GRPC_HANDLER{return &Handle{}}
