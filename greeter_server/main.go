/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Chen-Jin-yuan/GRPC-MyHelloWorld/helloworld"
	"github.com/Chen-Jin-yuan/grpc/consul"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	c  *consul.Client
	id string
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %s, %v", s.id, in.GetName())
	return &pb.HelloReply{Message: "Hello " + s.id + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloAgainRequest) (*pb.HelloAgainReply, error) {
	return &pb.HelloAgainReply{Message: "Hello again " + in.GetName(), DoubleNumber: in.GetNumber() * 2}, nil
}

func main() {
	flag.Parse()
	consulAddr := "127.0.0.1:8500"
	client, err := consul.NewClient(consulAddr)
	if err != nil {
		log.Fatalf("Got error while initializing Consul agent: %v", err)
	}
	num := 6
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		port := 50000 + i
		sid := strconv.Itoa(port)
		startServer(client, port, sid, &wg)
	}

	wg.Wait()
}

func startServer(client *consul.Client, port int, sid string, wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverIns := server{c: client, id: sid}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &serverIns)

	err = client.Register("helloServer", sid, "127.0.0.1", port)
	if err != nil {
		log.Fatalf("Got error while register service: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		wg.Done()
	}()

}
