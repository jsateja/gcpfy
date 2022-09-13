package meta

// code to initialize terraform

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
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

	return nil
}
