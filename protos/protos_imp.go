package protos

import(
	"github.com/ali2210/peers-peers/contacts"
	"context"
	"fmt"
)


// var(
// 	group contacts.Pal
// )

type Phonebook_Server struct {}

func(book *Phonebook_Server) AddContact(ctx context.Context, peers contacts.Friends)(contacts.Pal){
	 
	class := make([]contacts.Friends, 0)
	class = append(class, peers)
	 fmt.Println("Class:", class)
	 return contacts.Pal{}
}

// func NewSizeof(n int){group = make([]contacts.Pal, n)}

// func GetLen()int{return len(group)}
