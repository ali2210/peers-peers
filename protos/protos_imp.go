package protos



import(
	"github.com/ali2210/peers-peers/contacts"
)

var (
	lists_contact []contacts.List
)

type PhoneBookService interface {
	AddContact(contxt context.Context, peer contacts.Contact) (contacts.List)
	FindContact(contxt context.Context, info contacts.Phone)(contacts.List)
}

type PeerBook struct {}

func NewPeerBook() PhoneBookService{return &PeerBook{}}

func (book *PeerBook) AddContact(contxt context.Context, peer []contacts.Contact) (contacts.List){
	
	if len(peer) == 0{
		SizeOfListContacts(len(peer))
		return contacts.List{peer}
	}
	
	sizeOfListContacts(len(peer))
	
	for v := range getSize() {
		lists_contact = append(lists_contact, peer[v])
	}
	return lists_contact
}

func (book *PeerBook) FindContact(contxt context.Context, info contacts.Phone)(contacts.List){
	
	if GetSize() > 0{
			for v := range getListTxs() {
				if v == info.GetName(){
					return lists_contact[v]
				}
			}
	}
	return lists_contact[-1]
}


func sizeOfListContacts(n int){
	lists_contact := make([]contacts.Contact, n)
}

func getListTxs() []contacts.Contact{
	return lists_contact
}

func getSize() int{
	return len(lists_contact)-1
}


