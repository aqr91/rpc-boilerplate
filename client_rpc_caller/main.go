package main

import (
	"log"
	"net/rpc"
)

type RegisterArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReply struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := new(RegisterArgs)
	args.Username = "admin"
	args.Password = "admin"

	r := RegisterReply{}

	client.Call("Service.Login", args, &r)
	log.Println("reply: ", r)
}
