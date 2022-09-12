package main

import (
	asset "cloud.google.com/go/asset/apiv1"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

func main() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	req := &assetpb.SearchAllResourcesRequest{
		Scope: "projects/",
	}

	it := c.SearchAllResources(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		_ = resp
		fmt.Println(resp)
	}
}
