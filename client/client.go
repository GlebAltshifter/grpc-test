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
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/glebaltshifter/grpc-test/proto"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:8080"
	defaultDividend = 11
	defaultDivisor  = 4
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcTestClient(conn)

	// Contact the server and print out its response.
	dividend := defaultDividend
	divisor := defaultDivisor
	var callName string
	if len(os.Args) > 1 {
		var err error
		callName = os.Args[1]
		if ((callName == "GetQuotient") || (callName == "GetRemainder")) && (len(os.Args) > 3) {
			dividend, err = strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("cant read dividend: %v", err)
			}
			divisor, err = strconv.Atoi(os.Args[3])
			if err != nil {
				log.Fatalf("cant read divisor: %v", err)
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch callName {
	case "GetQuotient":
		{
			r, err := c.GetQuotient(ctx, &pb.DivisionPair{Dividend: int32(dividend), Divisor: int32(divisor)})
			if err != nil {
				log.Fatalf("could not get quotient: %v", err)
			}
			log.Printf("quotient: %d", r.Value)
		}
	case "GetRemainder":
		{
			r, err := c.GetRemainder(ctx, &pb.DivisionPair{Dividend: int32(dividend), Divisor: int32(divisor)})
			if err != nil {
				log.Fatalf("could not get remainder: %v", err)
			}
			log.Printf("remainder: %d", r.Value)
		}
	case "Lambs":
		{

		}
	}
}
