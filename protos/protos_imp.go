package protos

import(
	"github.com/ali2210/peers-peers/contacts"
	"context"
	//"fmt"
)


  var(
  	group contacts.Pal
	class []*contacts.Friends
 )

type Phonebook_Server struct {}

func(book *Phonebook_Server) AddContact(ctx context.Context, peers contacts.Friends)(contacts.Pal){

	class = make([]*contacts.Friends, 0)
	class = append(class, &peers)
	group.Palx = class
	 return group
}



