package address

import (
	"google.golang.org/grpc"
)
		

type (
	
	ClientPeersConnection struct{}

	ClientServiceGRPCOperation interface{
		OpenConnection(*grpc.ClientConn) ClientGRPCConnection
	}
	ClientGRPCConnection struct{
		Active *grpc.ClientConn
	}

	ServerPeersGRPCService struct{}
	
	PeerBookImpl interface{
		RegisterPeerService(*grpc.Server, *FriendsService)
	}
)

func NewContractClient() ClientServiceGRPCOperation  {
	return &ClientPeersConnection{}
}

func GRPCServerInit() PeerBookImpl  {
	return ServerPeersGRPCService{}
}

func (*ServerPeersGRPCService) RegisterPeerService(requestGRPC *grpcServer, userRequest *server.FriendsService) {
	Peer_ServiceDesc := grpc.ServiceDesc{
		ServiceName : "address.MyBook",
		HandlerType : (*server.PeersServer)(nil),
		Methods:[]grpc.MethodsDesc{
			{
				MethodName : AddPeers,
				Handler :_MyBook_AddPeers_Handler,
			},
			{
				MethodName : DeletePeers,
				Handler : _MyBook_DeletePeers_Handler,
			},
			{
				MethodName: DisplayPeers,
				Handler : _MyBook_DisplayPeers_Handler,
			},
		},
		Streams : []grpc.StreamDesc{},
		Metadata : "/pal/address/addresslist.proto",
	}
	requestGRPC.Register(Peer_ServiceDesc, userRequest)
}

func (*ClientPeersConnection) OpenConnection(connect *grpc.ClientConn) ClientGRPCConnection {
	return ClientGRPCConnection{Active: connect}
}