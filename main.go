package main

import (
	asset "cloud.google.com/go/asset/apiv1"
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"google.golang.org/api/iterator"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

func NewMeta() error {
	execPath := "/opt/homebrew/bin/terraform"

	workingDir := "/Users/jsate/GolandProjects/awesomeProject/gcpfy"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return fmt.Errorf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return fmt.Errorf("error running Init: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return fmt.Errorf("error running Show: %s", err)
	}
	fmt.Println(state.FormatVersion)

	ver, _, err := tf.Version(context.Background(), false)
	if err != nil {
		return fmt.Errorf("error running Version: %s", err)
	}
	fmt.Println(ver)
	return nil
}

func client() {
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

func main() {
	resp := NewMeta()
	fmt.Println(resp)
}
