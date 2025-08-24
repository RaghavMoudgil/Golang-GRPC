# Golang-GRPC
In this we are  going to create a GRPC project using golang


# basics 
There are 4 types of grpc API
1. Unary API (liek rest API server to client, client to server )-><-
2. server straming (server to client, sequence of msgs)<-
3. client streaming (client to server)->
4.Bi- directional straming(stream of data frokm both end client-><-server) this one is most common, this  asynconus both are working at the same time, it is not a queue but afterall that order of msgs is preserved

# setting up

after creating your protofile you need to generate go code for it 
    install protocol buffer for windows

    use protoc --version 

    use abvoe command to check proto is install correctly or not 
    if not change system variables 

# commad to conver proto file into go files 

protoc --go_out=. --go-grpc_out=. location of your proto file

with the above command two files will be created with pb.go extention 


# code in the serve to start the server 

# code in the client to start the client server 


    

 

