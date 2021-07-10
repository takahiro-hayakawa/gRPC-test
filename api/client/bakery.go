package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"pancake/maker/gen/api"
)

type Bakery struct {
}

func NewBakery() *Bakery {
	return &Bakery{}
}

func (*Bakery) bakePancake(c api.PancakeBakerServiceClient) {
	log.Println("Start bake")

	req := &api.BakeRequest{Menu: api.Pancake_CLASSIC}

	res, err := c.Bake(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling bake: %v", err)
	}
	log.Printf("Response from Server: %v\n", res)

}

func (*Bakery) report(c api.PancakeBakerServiceClient) {
	log.Println("Start report")

	req := &api.ReportRequest{}

	res, err := c.Report(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling report: %v", err)
	}
	log.Printf("Response from Server: %v\n", res)
}

func main() {
	var opts []grpc.DialOption
	tls := false
	if tls {
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := api.NewPancakeBakerServiceClient(cc)
	bakery := NewBakery()
	bakery.bakePancake(c)
	bakery.report(c)
}
