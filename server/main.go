package main


import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/ali2210/pal/address"
)

const(
	port string = ":5000"
	handshake string = "tcp"
)

var(
	count_PEERS int32 = 0
)


type Machine struct{
	pb.UnimplementedContactServer
}

func (*Machine) AddPeers(ctx Context.Context, peer *pb.Peers)(*pb.Contacts, error ) {
	 if peer.GetName() == "" && peer.GetPwallet() == ""{
		 return pb.Contacts{}, errors.New("No user generate")
	 }
	 peer.Uid = getPeersCount()
	 return &pb.Contacts{friend:peer}, nil
}

func (*Machine) DeletePeers(ctx Context.Context, peer *pb.Peers) (*pb.Contacts, error)  {
	contact := &pb.Contacts{}
	List := contact.GetFriend()
	
	for i := range List {
		if List[i] == peer{
			return List[i]{
				friend : &peer{}
			}, nil
		}
	}
}

func (*Machine) DisplayPeers(ctx Context.Context, c *pb.Contacts) (*pb.Contacts, error)  {
	return c.GetFriend(), nil
}


func main() {
	connection , err := listen(handshake, port)
	if err != nil{
		log.Fatal("Connection handshake ", err)
		return 
	}
	grpcServer := grpc.NewServer()
}


func getPeersCount()  {
	return count_PEERS+1
}

