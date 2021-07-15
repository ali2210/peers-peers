# peers-peers
Project must be inside go package src 


    ## Steps for install protobuf 
       
       ### Go version +1.6
       
       ### $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
       
       ### $ export PATH="$PATH:$(go env GOPATH)/bin"
       
       ### $ protoc --go_out=paths=source_relative:. contacts/phonebook.proto 
       
       ## code-generation: https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go
      
      Protobuf-compiler [ok]
