package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"pancake/maker/gen/api"
)

type Bakery struct {
}

func NewBakery() *Bakery {
	return &Bakery{}
}

func (*Bakery) pbMenu(menu string) api.Pancake_Menu {
	switch menu {
	case "classic":
		return api.Pancake_CLASSIC
	case "banana_and_whip":
		return api.Pancake_BANANA_AND_WHIP
	case "bacon_and_cheese":
		return api.Pancake_BANANA_AND_WHIP
	case "mix_berry":
		return api.Pancake_MINI_BERRY
	case "baked_marshmallow":
		return api.Pancake_BAKED_MARSHMALLOW
	case "spicy_curry":
		return api.Pancake_SPICY_CURRY
	default:
		return api.Pancake_UNKNOWN
	}
}

func (bakery *Bakery) bakePancake(menu string, c api.PancakeBakerServiceClient) {
	log.Println("Start bake")

	req := &api.BakeRequest{Menu: bakery.pbMenu(menu)}

	md := metadata.New(map[string]string{"authorization": "bearer hi/mi/tsu"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.Bake(ctx, req)
	if err != nil {
		log.Fatalf("error while calling bake: %v", err)
	}
	log.Printf("Response from Server: %v\n", res)

}

func (*Bakery) report(c api.PancakeBakerServiceClient) {
	log.Println("Start report")

	req := &api.ReportRequest{}

	md := metadata.New(map[string]string{"authorization": "bearer hi/mi/tsu"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.Report(ctx, req)
	if err != nil {
		log.Fatalf("error while calling report: %v", err)
	}
	log.Printf("Response from Server: %v\n", res)
}

func main() {
	var opts []grpc.DialOption

	// セキュア通信設定
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
	bakery.bakePancake("classic", c)
	bakery.bakePancake("banana_and_whip", c)
	bakery.report(c)
}
