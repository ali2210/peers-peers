package contacts

import (
	context "context"
	grpc "google.golang.org/grpc"
	"errors"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"fmt"
)

const _ = grpc.SupportPackageIsVersion7

type Phonebook_Client interface {
	AddContact(ctx context.Context, peer *Friends, opts ...grpc.CallOption)(*Pal)
}

type peers_Node struct { grpc_interface grpc.ClientConnInterface}

func NewClientServiceInit(this_interface grpc.ClientConnInterface) Phonebook_Client  { return &peers_Node{this_interface}}

func (p *peers_Node) AddContact(ctx context.Context, peer *Friends, opts ...grpc.CallOption)(*Pal)  {
	service_add_contact := new(Pal)
	err := p.grpc_interface.Invoke(ctx, "/contacts.Phonebook/AddContact", peer, service_add_contact, opts...)
	if err != nil{
		class := make([]*Friends, 0)
		class = append(class, peer)
		service_add_contact.Palx = class
		return service_add_contact
	}
	return service_add_contact
}

type PhonebookServers interface {
	AddContact(ctx context.Context, peer *Friends)(*Pal)	
	mustUnimplementedService()
}

type UnimplementedServer struct {}

func (UnimplementedServer) AddContact(ctx context.Context, peers *Friends)(*Pal)  {
	out_service := new(Pal)
	fmt.Println(status.Errorf(codes.Unimplemented, "method Add-Contact not implemented"))
	return out_service
}
func (UnimplementedServer) mustUnimplementedService(){}
type UnSafeBehaviour interface {
	mustUnimplementedService()
}

func RegisterPhoneService(service grpc.ServiceRegistrar, srv PhonebookServers){
	service.RegisterService(&Phone_AddServiceDesc, srv)
}

func _PhoneBook_Service(grpc_server interface{}, ctx context.Context, des func (interface{}) error, interceptor grpc.UnaryServerInterceptor)(interface{}, error) {
	service_add_contact := new(Friends)
	err := des(service_add_contact); if err != nil {
		return nil, err
	}

	if interceptor == nil {
		grpc_server.(PhonebookServers).AddContact(ctx, service_add_contact)
	}
	unary := &grpc.UnaryServerInfo{
		Server : grpc_server,
		FullMethod: "/contacts.Phonebook/AddContact",
	}
	handle := func(ctx context.Context, request interface{})(interface{}, error){
		friend := grpc_server.(PhonebookServers).AddContact(ctx, request.(*Friends))
		return friend, errors.New("Request Fail")
	}
	return interceptor(ctx,service_add_contact,unary, handle)
}

var Phone_AddServiceDesc = grpc.ServiceDesc{
	ServiceName : "contacts.Phonebook",
	HandlerType : (*PhonebookServers)(nil),
	Methods:  []grpc.MethodDesc{
		{
			MethodName: "AddContact",
			Handler:    _PhoneBook_Service,
		},
	},
	Streams : []grpc.StreamDesc{},
	Metadata : "github.com/ali2210/peers-peers/contacts/users.proto",
}
