# peers-peers
    Project must be inside go package src 

    ## Steps for install protobuf 
       
       ### Go version +1.6
       
       ### $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
       
       ### $ export PATH="$PATH:$(go env GOPATH)/bin"
       
       ### $ protoc --go_out=paths=source_relative:. contacts/phonebook.proto 
       
       ## code-generation: https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go
      
      Protobuf-compiler [ok]

    ## Project-Goal
     
     1. Execute program as protobuf 
          
          ## go run main.go

     2. Execute grpc-proto together 

         ## forked project / download project inside go/src pkg
         
         ## cd server 
         
         ## go run main.go 

         ## open new terminal ctrl+T (LINUX)

         ## cd client

         ## go run main.go

         ## enjoy 

          
