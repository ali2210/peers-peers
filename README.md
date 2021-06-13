# peers-peers
Project must be inside go package src 


##. Steps for install protobuf 
 
  ###. Go version +1.6
  ###. $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ###. $ export PATH="$PATH:$(go env GOPATH)/bin"

Now compile your code 
 ##. $ protoc -I= . --go_out=. /$PATH_WHERE_PROTO_FILE
