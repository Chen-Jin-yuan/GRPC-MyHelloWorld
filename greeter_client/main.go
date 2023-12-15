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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	pb "github.com/Chen-Jin-yuan/GRPC-MyHelloWorld/helloworld"
	"github.com/Chen-Jin-yuan/grpc/consul"
	"github.com/Chen-Jin-yuan/grpc/dialer"
)

const (
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	//consulAddr := "127.0.0.1:8500"
	consulAddr := "consul:8500"
	svcname := "helloServer"
	consulClient, err := consul.NewClient(consulAddr)

	conn, err := dialer.Dial(
		svcname,
		// 路径从项目根路径开始，
		dialer.WithBalancerBF(consulClient, "./greeter_client/config.json", 10001),
		dialer.WithStatsHandlerBF(),
	)
	//conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	md := metadata.Pairs("request-type", "v1")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	for {
		r1, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r1.GetMessage())

		time.Sleep(1e7)
	}

}
